package services

import (
	"mime/multipart"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"
)

type ProfileService interface {
	GetUserByID(userID string) (*models.User, error)
	UpdateProfile(userID string, req dto.UpdateProfileRequest) error
	UpdateAvatar(userID string, file *multipart.FileHeader) (string, error)
	GetUserTransactions(userID string, page, limit int) (*dto.TransactionListResponse, error)
}

type profileService struct {
	repo repositories.ProfileRepository
}

func NewProfileService(repo repositories.ProfileRepository) ProfileService {
	return &profileService{repo}
}

func (s *profileService) GetUserByID(userID string) (*models.User, error) {
	return s.repo.GetUserByID(userID)
}
func (s *profileService) UpdateProfile(userID string, req dto.UpdateProfileRequest) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.Profile.Fullname = req.Fullname
	user.Profile.Gender = req.Gender
	user.Profile.Phone = req.Phone
	user.Profile.Bio = req.Bio
	if req.Birthday != "" {
		birthday, err := time.Parse("2006-01-02", req.Birthday)
		if err == nil {
			user.Profile.Birthday = &birthday
		}
	}

	return s.repo.UpdateUser(user)
}

func (s *profileService) UpdateAvatar(userID string, file *multipart.FileHeader) (string, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return "", err
	}

	if file == nil {
		return "", nil
	}

	if err := utils.ValidateImageFile(file); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	newAvatarURL, err := utils.UploadToCloudinary(src)
	if err != nil {
		return "", err
	}

	if user.Profile.Avatar != "" && user.Profile.Avatar != newAvatarURL && !isDiceBear(user.Profile.Avatar) {
		_ = utils.DeleteFromCloudinary(user.Profile.Avatar)
	}

	user.Profile.Avatar = newAvatarURL
	err = s.repo.UpdateUser(user)
	return newAvatarURL, err
}

func isDiceBear(url string) bool {
	return url != "" && (len(url) > 0 && (url[:30] == "https://api.dicebear.com" || url[:31] == "https://avatars.dicebear.com"))
}

func (s *profileService) GetUserTransactions(userID string, page, limit int) (*dto.TransactionListResponse, error) {
	offset := (page - 1) * limit

	payments, total, err := s.repo.GetUserTransactions(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	var transactions []dto.TransactionResponse
	for _, p := range payments {
		transactions = append(transactions, dto.TransactionResponse{
			ID:            p.ID.String(),
			PackageID:     p.PackageID.String(),
			PackageName:   p.Package.Name,
			PaymentMethod: p.PaymentMethod,
			Status:        p.Status,
			Price:         p.Package.Price,
			PaidAt:        p.PaidAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &dto.TransactionListResponse{
		Transactions: transactions,
		Total:        total,
		Page:         page,
		Limit:        limit,
	}, nil
}
