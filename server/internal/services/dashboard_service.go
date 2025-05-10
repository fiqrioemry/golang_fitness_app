package services

import (
	"server/internal/dto"
	"server/internal/repositories"
)

type DashboardService interface {
	GetSummary() (*dto.DashboardSummaryResponse, error)
	GetRevenueStats(filter dto.RevenueStatRequest) (*dto.RevenueStatResponse, error)
}

type dashboardService struct {
	repo repositories.DashboardRepository
}

func NewDashboardService(repo repositories.DashboardRepository) DashboardService {
	return &dashboardService{repo}
}

func (s *dashboardService) GetSummary() (*dto.DashboardSummaryResponse, error) {
	users, _ := s.repo.CountUsers()
	instructors, _ := s.repo.CountInstructors()
	classes, _ := s.repo.CountClasses()
	bookings, _ := s.repo.CountBookings()
	payments, _ := s.repo.CountPayments()
	revenue, _ := s.repo.SumRevenue()
	packages, _ := s.repo.CountActivePackages()
	attended, _ := s.repo.CountAttendanceByStatus("attended")
	absent, _ := s.repo.CountAttendanceByStatus("absent")

	return &dto.DashboardSummaryResponse{
		TotalUsers:         users,
		TotalInstructors:   instructors,
		TotalClasses:       classes,
		TotalBookings:      bookings,
		TotalPayments:      payments,
		TotalRevenue:       revenue,
		ActivePackages:     packages,
		TotalAttendance:    attended + absent,
		AttendedAttendance: attended,
		AbsentAttendance:   absent,
	}, nil
}

func (s *dashboardService) GetRevenueStats(filter dto.RevenueStatRequest) (*dto.RevenueStatResponse, error) {
	return s.repo.GetRevenueStats(filter)
}
