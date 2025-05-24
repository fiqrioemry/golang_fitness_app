// internal/services/user_service.go
package services

import (
	"server/internal/dto"
	"server/internal/repositories"
)

type UserService interface {
	GetAllUsers(params dto.UserQueryParam) ([]dto.UserListResponse, *dto.PaginationResponse, error)
	GetUserDetail(id string) (*dto.UserDetailResponse, error)
	GetUserStats() (*dto.UserStatsResponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAllUsers(params dto.UserQueryParam) ([]dto.UserListResponse, *dto.PaginationResponse, error) {

	users, total, err := s.repo.FindAllUsers(params)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.UserListResponse
	for _, u := range users {
		result = append(result, dto.UserListResponse{
			ID:       u.ID.String(),
			Email:    u.Email,
			Role:     u.Role,
			Fullname: u.Profile.Fullname,
			Phone:    u.Profile.Phone,
			Avatar:   u.Profile.Avatar,
			JoinedAt: u.CreatedAt.Format("2006-01-02"),
		})
	}
	totalPages := int((total + int64(params.Limit) - 1) / int64(params.Limit))

	pagination := &dto.PaginationResponse{
		Page:       params.Page,
		Limit:      params.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}

	return result, pagination, nil

}

func (s *userService) GetUserDetail(id string) (*dto.UserDetailResponse, error) {
	u, err := s.repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	res := &dto.UserDetailResponse{
		ID:       u.ID.String(),
		Email:    u.Email,
		Role:     u.Role,
		Fullname: u.Profile.Fullname,
		Phone:    u.Profile.Phone,
		Avatar:   u.Profile.Avatar,
		Gender:   u.Profile.Gender,
		Bio:      u.Profile.Bio,
		JoinedAt: u.CreatedAt.Format("2006-01-02"),
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
