// internal/services/user_service.go
package services

import (
	"server/internal/dto"
	"server/internal/repositories"
	"time"
)

type UserService interface {
	GetAllUsers(params dto.UserQueryParam) ([]dto.UserListResponse, int64, error)
	GetUserDetail(id string) (*dto.UserDetailResponse, error)
	GetUserStats() (*dto.UserStatsResponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAllUsers(params dto.UserQueryParam) ([]dto.UserListResponse, int64, error) {
	offset := (params.Page - 1) * params.Limit
	users, total, err := s.repo.FindAllUsers(params.Q, params.Role, params.Sort, params.Limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var result []dto.UserListResponse
	for _, u := range users {
		result = append(result, dto.UserListResponse{
			ID:        u.ID.String(),
			Email:     u.Email,
			Role:      u.Role,
			Fullname:  u.Profile.Fullname,
			Phone:     u.Profile.Phone,
			Avatar:    u.Profile.Avatar,
			CreatedAt: u.CreatedAt.Format(time.RFC3339),
		})
	}
	return result, total, nil
}

func (s *userService) GetUserDetail(id string) (*dto.UserDetailResponse, error) {
	u, err := s.repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	res := &dto.UserDetailResponse{
		ID:        u.ID.String(),
		Email:     u.Email,
		Role:      u.Role,
		Fullname:  u.Profile.Fullname,
		Phone:     u.Profile.Phone,
		Avatar:    u.Profile.Avatar,
		Gender:    u.Profile.Gender,
		Bio:       u.Profile.Bio,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
	if u.Profile.Birthday != nil {
		res.Birthday = u.Profile.Birthday.Format("2006-01-02")
	}
	return res, nil
}

func (s *userService) GetUserStats() (*dto.UserStatsResponse, error) {
	total, customers, instructors, admins, newMonth, err := s.repo.GetUserStats()
	if err != nil {
		return nil, err
	}
	return &dto.UserStatsResponse{
		Total:        total,
		Customers:    customers,
		Instructors:  instructors,
		Admins:       admins,
		NewThisMonth: newMonth,
	}, nil
}
