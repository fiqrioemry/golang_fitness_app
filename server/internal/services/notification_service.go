package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type NotificationService interface {
	GetSettingsByUser(userID string) ([]dto.NotificationSettingResponse, error)
	UpdateSetting(userID string, req dto.UpdateNotificationSettingRequest) error
	CreateNotification(input dto.CreateNotificationRequest) error
	GetUnreadNotifications(userID string) ([]models.Notification, error)
	MarkAsRead(userID, notifID string) error
}

type notificationService struct {
	repo repositories.NotificationRepository
}

func NewNotificationService(repo repositories.NotificationRepository) NotificationService {
	return &notificationService{repo}
}

func (s *notificationService) GetSettingsByUser(userID string) ([]dto.NotificationSettingResponse, error) {
	settings, err := s.repo.GetNotificationSettingsByUser(uuid.MustParse(userID))
	if err != nil {
		return nil, err
	}

	var result []dto.NotificationSettingResponse
	for _, s := range settings {
		result = append(result, dto.NotificationSettingResponse{
			TypeID:  s.NotificationTypeID.String(),
			Code:    s.NotificationType.Code,
			Title:   s.NotificationType.Title,
			Channel: s.Channel,
			Enabled: s.Enabled,
		})
	}

	return result, nil
}

func (s *notificationService) UpdateSetting(userID string, req dto.UpdateNotificationSettingRequest) error {
	typeID, err := uuid.Parse(req.TypeID)
	if err != nil {
		return err
	}

	setting, err := s.repo.FindSetting(uuid.MustParse(userID), typeID, req.Channel)
	if err != nil {
		return err
	}

	setting.Enabled = req.Enabled
	return s.repo.UpdateNotificationSetting(setting)
}

func (s *notificationService) CreateNotification(input dto.CreateNotificationRequest) error {
	notif := models.Notification{
		ID:       uuid.New(),
		UserID:   uuid.MustParse(input.UserID),
		TypeCode: input.TypeCode,
		Title:    input.Title,
		Message:  input.Message,
		Channel:  input.Channel,
	}

	return s.repo.CreateNotification(&notif)
}

func (s *notificationService) GetUnreadNotifications(userID string) ([]models.Notification, error) {
	return s.repo.GetUserUnreadNotifications(uuid.MustParse(userID))
}

func (s *notificationService) MarkAsRead(userID, notifID string) error {
	return s.repo.MarkNotificationRead(uuid.MustParse(userID), uuid.MustParse(userID))
}
