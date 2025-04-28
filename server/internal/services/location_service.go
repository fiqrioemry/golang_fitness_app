package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type LocationService interface {
	CreateLocation(req dto.CreateLocationRequest) error
	UpdateLocation(id string, req dto.UpdateLocationRequest) error
	DeleteLocation(id string) error
	GetAllLocations() ([]dto.LocationResponse, error)
	GetLocationByID(id string) (*dto.LocationResponse, error)
}

type locationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) LocationService {
	return &locationService{repo}
}

func (s *locationService) CreateLocation(req dto.CreateLocationRequest) error {
	location := models.Location{
		ID:          uuid.New(),
		Name:        req.Name,
		Address:     req.Address,
		GeoLocation: req.GeoLocation,
	}
	return s.repo.CreateLocation(&location)
}

func (s *locationService) UpdateLocation(id string, req dto.UpdateLocationRequest) error {
	location, err := s.repo.GetLocationByID(id)
	if err != nil {
		return err
	}

	location.Name = req.Name
	location.Address = req.Address
	location.GeoLocation = req.GeoLocation

	return s.repo.UpdateLocation(location)
}

func (s *locationService) DeleteLocation(id string) error {
	return s.repo.DeleteLocation(id)
}

func (s *locationService) GetAllLocations() ([]dto.LocationResponse, error) {
	locations, err := s.repo.GetAllLocations()
	if err != nil {
		return nil, err
	}

	var result []dto.LocationResponse
	for _, l := range locations {
		result = append(result, dto.LocationResponse{
			ID:          l.ID.String(),
			Name:        l.Name,
			Address:     l.Address,
			GeoLocation: l.GeoLocation,
		})
	}
	return result, nil
}

func (s *locationService) GetLocationByID(id string) (*dto.LocationResponse, error) {
	location, err := s.repo.GetLocationByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.LocationResponse{
		ID:          location.ID.String(),
		Name:        location.Name,
		Address:     location.Address,
		GeoLocation: location.GeoLocation,
	}, nil
}
