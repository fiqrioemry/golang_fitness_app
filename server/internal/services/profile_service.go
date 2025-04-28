package services

import (
	"server/internal/models"
	"server/internal/repositories"
)

type ProfileService interface {
	GetUserProfile(userID string) (*models.User, error)
}
type profileService struct {
	profileRepo repositories.ProfileRepository
	authService *authService
}

func NewProfileService(profileRepo repositories.ProfileRepository, authService *AuthService) *ProfileService {
	return &profileService{
		profileRepo: profileRepo,
		authService: authService,
	}
}
func (s *authService) GetUserProfile(userID string) (*models.User, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
