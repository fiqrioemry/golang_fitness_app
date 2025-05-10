package repositories

import (
	"server/internal/dto"
	"server/internal/models"
	"time"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	CountUsers() (int, error)
	CountInstructors() (int, error)
	CountClasses() (int, error)
	CountBookings() (int, error)
	CountPayments() (int, error)
	SumRevenue() (float64, error)
	CountActivePackages() (int, error)
	GetRevenueStats(filter dto.RevenueStatRequest) (*dto.RevenueStatResponse, error)
	CountAttendanceByStatus(status string) (int, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db}
}

func (r *dashboardRepository) CountUsers() (int, error) {
	var count int64
	err := r.db.Model(&models.User{}).Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) CountInstructors() (int, error) {
	var count int64
	err := r.db.Model(&models.Instructor{}).Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) CountClasses() (int, error) {
	var count int64
	err := r.db.Model(&models.Class{}).Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) CountBookings() (int, error) {
	var count int64
	err := r.db.Model(&models.Booking{}).Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) CountPayments() (int, error) {
	var count int64
	err := r.db.Model(&models.Payment{}).Where("status = ?", "success").Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) SumRevenue() (float64, error) {
	var total float64
	err := r.db.Model(&models.Payment{}).Where("status = ?", "success").Select("SUM(total)").Scan(&total).Error
	return total, err
}

func (r *dashboardRepository) CountActivePackages() (int, error) {
	var count int64
	err := r.db.Model(&models.Package{}).Where("is_active = true").Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) CountAttendanceByStatus(status string) (int, error) {
	var count int64
	err := r.db.Model(&models.Attendance{}).Where("status = ?", status).Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) GetRevenueStats(filter dto.RevenueStatRequest) (*dto.RevenueStatResponse, error) {
	var result dto.RevenueStatResponse

	var startDate time.Time
	now := time.Now()

	switch filter.Range {
	case "daily":
		startDate = now.Truncate(24 * time.Hour)
	case "monthly":
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	case "yearly":
		startDate = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	default:
		startDate = time.Time{} // All time
	}

	query := r.db.Model(&models.Payment{}).Where("paid_at >= ?", startDate)

	if err := query.
		Where("status = ?", "success").
		Select("COALESCE(SUM(total), 0) as total_revenue, COUNT(*) as total_success").
		Scan(&result).Error; err != nil {
		return nil, err
	}

	var pending, failed int64

	if err := query.
		Where("status = ?", "pending").
		Count(&pending).Error; err != nil {
		return nil, err
	}

	if err := query.
		Where("status = ?", "failed").
		Count(&failed).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
