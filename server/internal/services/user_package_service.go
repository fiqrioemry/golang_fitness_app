package services

import (
	"server/internal/dto"
	"server/internal/repositories"
)

type UserPackageService interface {
	GetUserPackages(userID string) ([]dto.UserPackageResponse, error)
}

type userPackageService struct {
	repo repositories.UserPackageRepository
}

func NewUserPackageService(repo repositories.UserPackageRepository) UserPackageService {
	return &userPackageService{repo}
}

func (s *userPackageService) GetUserPackages(userID string) ([]dto.UserPackageResponse, error) {
	userPackages, err := s.repo.GetUserPackagesByUserID(userID)
	if err != nil {
		return nil, err
	}

	var result []dto.UserPackageResponse
	for _, up := range userPackages {
		expiredAt := ""
		if up.ExpiredAt != nil {
			expiredAt = up.ExpiredAt.Format("2006-01-02")
		}

		result = append(result, dto.UserPackageResponse{
			ID:              up.ID.String(),
			PackageName:     "", // package name bisa diisi kalau mau preload, sekarang kosong dulu
			RemainingCredit: up.RemainingCredit,
			ExpiredAt:       expiredAt,
			PurchasedAt:     up.PurchasedAt.Format("2006-01-02"),
		})
	}

	return result, nil
}
