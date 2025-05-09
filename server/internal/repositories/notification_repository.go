package repositories

import (
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	MarkNotificationRead(userID, notifID uuid.UUID) error
	InsertNotifications(notifs []models.Notification) error
	CreateNotification(notification *models.Notification) error
	GetAllNotificationTypes() ([]models.NotificationType, error)
	UpdateNotificationSetting(setting *models.NotificationSetting) error
	CreateNotificationSetting(setting *models.NotificationSetting) error
	GetUserNotifications(userID uuid.UUID) ([]models.Notification, error)
	GetNotificationSettingsByUser(userID uuid.UUID) ([]models.NotificationSetting, error)
	FindSetting(userID, typeID uuid.UUID, channel string) (*models.NotificationSetting, error)
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) GetUserNotifications(userID uuid.UUID) ([]models.Notification, error) {
	var notifs []models.Notification
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&notifs).Error
	return notifs, err
}

func (r *notificationRepository) GetAllNotificationTypes() ([]models.NotificationType, error) {
	var types []models.NotificationType
	err := r.db.Find(&types).Error
	return types, err
}

func (r *notificationRepository) GetNotificationSettingsByUser(userID uuid.UUID) ([]models.NotificationSetting, error) {
	var settings []models.NotificationSetting
	err := r.db.
		Preload("NotificationType").
		Where("user_id = ?", userID).
		Find(&settings).Error
	return settings, err
}

func (r *notificationRepository) FindSetting(userID, typeID uuid.UUID, channel string) (*models.NotificationSetting, error) {
	var setting models.NotificationSetting
	err := r.db.
		Where("user_id = ? AND notification_type_id = ? AND channel = ?", userID, typeID, channel).
		First(&setting).Error
	return &setting, err
}

func (r *notificationRepository) UpdateNotificationSetting(setting *models.NotificationSetting) error {
	return r.db.Save(setting).Error
}

func (r *notificationRepository) InsertNotifications(notifs []models.Notification) error {
	return r.db.Create(&notifs).Error
}

func (r *notificationRepository) CreateNotification(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) MarkNotificationRead(userID, notifID uuid.UUID) error {
	return r.db.Model(&models.Notification{}).
		Where("user_id = ? AND id = ?", userID, notifID).
		Update("is_read", true).Error
}

func (r *notificationRepository) CreateNotificationSetting(setting *models.NotificationSetting) error {
	return r.db.Create(setting).Error
}
