package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/idtoken"
)

type AuthService interface {
	SendOTP(email string) error
	Logout(refreshToken string) error
	VerifyOTP(email, otp string) error
	GetUserProfile(userID string) (*models.User, error)
	Login(req *dto.LoginRequest) (*dto.AuthResponse, error)
	GoogleSignIn(idToken string) (*dto.AuthResponse, error)
	Register(req *dto.RegisterRequest) (*dto.AuthResponse, error)
	RefreshToken(refreshToken string) (*dto.AuthResponse, error)
}

type authService struct {
	repo             repositories.AuthRepository
	notificationRepo repositories.NotificationRepository
}

func NewAuthService(repo repositories.AuthRepository, notificationRepo repositories.NotificationRepository) AuthService {
	return &authService{repo: repo, notificationRepo: notificationRepo}
}

func (s *authService) SendOTP(email string) error {
	user, err := s.repo.GetUserByEmail(email)
	if err == nil {
		if user.Password == "-" {
			return errors.New("email already registered via Google Sign-In")
		}
		return errors.New("email already registered")
	}
	subject := "One-Time Password (OTP) xxxxx"
	otp := utils.GenerateOTP(6)
	body := fmt.Sprintf("Your OTP code is %s", otp)

	err = utils.SendEmail(subject, email, otp, body)
	if err != nil {
		return errors.New("failed to send email")
	}

	err = config.RedisClient.Set(config.Ctx, "otp:"+email, otp, 5*time.Minute).Err()
	return err
}

func (s *authService) Logout(refreshToken string) error {
	return s.repo.DeleteRefreshToken(refreshToken)
}

func (s *authService) VerifyOTP(email, otp string) error {
	savedOtp, err := config.RedisClient.Get(config.Ctx, "otp:"+email).Result()
	if err != nil {
		return errors.New("otp expired or invalid")
	}

	if savedOtp != otp {
		return errors.New("invalid OTP code")
	}

	config.RedisClient.Del(config.Ctx, "otp:"+email)
	return nil
}

func (s *authService) GetUserProfile(userID string) (*models.User, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	tokenModel := &models.Token{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiredAt: time.Now().Add(7 * 24 * time.Hour),
	}
	if err := s.repo.StoreRefreshToken(tokenModel); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	profile := models.Profile{
		Fullname: req.Fullname,
		Avatar:   utils.RandomUserAvatar(),
	}

	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "customer",
		Profile:  profile,
	}

	if err := s.repo.CreateUser(&user); err != nil {
		return nil, err
	}

	userID := user.ID.String()

	accessToken, err := utils.GenerateAccessToken(userID, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	tokenModel := &models.Token{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiredAt: time.Now().Add(7 * 24 * time.Hour),
	}
	if err := s.repo.StoreRefreshToken(tokenModel); err != nil {
		return nil, err
	}

	s.generateDefaultSettingsForUser(user.ID)

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) RefreshToken(refreshToken string) (*dto.AuthResponse, error) {

	_, err := utils.DecodeRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	tokenModel, err := s.repo.FindRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("refresh token not found")
	}

	if tokenModel.ExpiredAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	user, err := s.repo.GetUserByID(tokenModel.UserID.String())
	if err != nil {
		return nil, errors.New("user not found")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	if err := s.repo.DeleteRefreshToken(refreshToken); err != nil {
		return nil, err
	}

	newToken := &models.Token{
		UserID:    user.ID,
		Token:     newRefreshToken,
		ExpiredAt: time.Now().Add(7 * 24 * time.Hour),
	}
	if err := s.repo.StoreRefreshToken(newToken); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *authService) generateDefaultSettingsForUser(userID uuid.UUID) {
	notifTypes, _ := s.notificationRepo.GetAllNotificationTypes()

	for _, nt := range notifTypes {
		for _, channel := range []string{"email", "browser"} {
			setting := models.NotificationSetting{
				ID:                 uuid.New(),
				UserID:             userID,
				NotificationTypeID: nt.ID,
				Channel:            channel,
				Enabled:            nt.DefaultEnabled,
			}
			_ = s.notificationRepo.CreateNotificationSetting(&setting)
		}
	}
}

func (s *authService) GoogleSignIn(idToken string) (*dto.AuthResponse, error) {
	payload, err := idtoken.Validate(context.Background(), idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return nil, errors.New("invalid Google ID token")
	}

	email, ok := payload.Claims["email"].(string)
	if !ok || email == "" {
		return nil, errors.New("email not found in token")
	}
	name, _ := payload.Claims["name"].(string)
	picture, _ := payload.Claims["picture"].(string)

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		user = &models.User{
			Email:    email,
			Password: "-",
			Role:     "customer",
			Profile: models.Profile{
				Fullname: name,
				Avatar:   picture,
			},
		}

		if err := s.repo.CreateUser(user); err != nil {
			return nil, err
		}

		if user.ID == uuid.Nil {
			return nil, errors.New("failed to assign UUID to user")
		}
		fmt.Println("✅ User created with ID:", user.ID)

		s.generateDefaultSettingsForUser(user.ID)
	}

	fmt.Println("➡️ Login Google untuk user ID:", user.ID)

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	tokenModel := &models.Token{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiredAt: time.Now().Add(7 * 24 * time.Hour),
	}

	// ⛏️ Cek user ID sebelum simpan token
	if tokenModel.UserID == uuid.Nil {
		return nil, errors.New("user ID kosong saat menyimpan token")
	}
	fmt.Println("✅ Simpan refresh token untuk user ID:", tokenModel.UserID)

	if err := s.repo.StoreRefreshToken(tokenModel); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
