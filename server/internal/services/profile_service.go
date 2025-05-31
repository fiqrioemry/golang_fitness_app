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
	UpdateAvatar(userID string, file *multipart.FileHeader) error
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

	user.Profile.Bio = req.Bio
	user.Profile.Phone = req.Phone
	user.Profile.Gender = req.Gender
	user.Profile.Fullname = req.Fullname
	if req.Birthday != "" {
		birthday, err := time.Parse("2006-01-02", req.Birthday)
		if err == nil {
			user.Profile.Birthday = &birthday
		}
	}

	return s.repo.UpdateUser(user)
}

func (s *profileService) UpdateAvatar(userID string, file *multipart.FileHeader) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if file == nil {
		return nil
	}

	if err := utils.ValidateImageFile(file); err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	newAvatarURL, err := utils.UploadToCloudinary(src)
	if err != nil {
		return err
	}

	if user.Profile.Avatar != "" && user.Profile.Avatar != newAvatarURL && !isDiceBear(user.Profile.Avatar) {
		_ = utils.DeleteFromCloudinary(user.Profile.Avatar)
	}

	user.Profile.Avatar = newAvatarURL
	err = s.repo.UpdateUser(user)
	return err
}

func isDiceBear(url string) bool {
	return url != "" && (len(url) > 0 && (url[:30] == "https://api.dicebear.com" || url[:31] == "https://avatars.dicebear.com"))
}
