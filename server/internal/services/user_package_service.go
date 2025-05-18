package services

import (
	"server/internal/dto"
	"server/internal/repositories"
	"time"
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
		var (
			expiredAt     string
			expiredInDays int
		)

		if up.ExpiredAt != nil {
			expiredAt = up.ExpiredAt.Format("2006-01-02")
			expiredInDays = int(max(0, int(time.Until(*up.ExpiredAt).Hours()/24)))
		}

		result = append(result, dto.UserPackageResponse{
			ID:              up.ID.String(),
			PackageID:       up.PackageID.String(),
			PackageName:     up.PackageName,
			RemainingCredit: up.RemainingCredit,
			ExpiredAt:       expiredAt,
			ExpiredInDays:   expiredInDays,
			PurchasedAt:     up.PurchasedAt.Format("2006-01-02"),
		})
	}

	return result, nil
}
