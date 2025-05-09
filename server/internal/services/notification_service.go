package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type NotificationService interface {
	MarkAllAsRead(userID string) error
	CreateNotification(input dto.CreateNotificationRequest) error
	SendPromoNotification(req dto.SendPromoNotificationRequest) error
	GetAllNotifications(userID string) ([]dto.NotificationResponse, error)
	UpdateSetting(userID string, req dto.UpdateNotificationSettingRequest) error
	GetSettingsByUser(userID string) ([]dto.NotificationSettingResponse, error)
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

func (s *notificationService) GetAllNotifications(userID string) ([]dto.NotificationResponse, error) {
	notifs, err := s.repo.GetAllBrowserNotifications(uuid.MustParse(userID))
	if err != nil {
		return nil, err
	}

	var result []dto.NotificationResponse
	for _, n := range notifs {
		result = append(result, dto.NotificationResponse{
			ID:        n.ID.String(),
			TypeCode:  n.TypeCode,
			Title:     n.Title,
			Message:   n.Message,
			Channel:   n.Channel,
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt.Format(time.RFC3339),
		})
	}

	return result, nil
}

func (s *notificationService) MarkAllAsRead(userID string) error {
	return s.repo.MarkAllNotificationsRead(uuid.MustParse(userID))
}

func (s *notificationService) SendPromoNotification(req dto.SendPromoNotificationRequest) error {
	settings, err := s.repo.GetUsersWithEnabledPromoNotifications()
	if err != nil {
		return err
	}

	var notifs []models.Notification
	for _, setting := range settings {
		notifs = append(notifs, models.Notification{
			ID:       uuid.New(),
			UserID:   setting.UserID,
			TypeCode: "promo_offer",
			Title:    req.Title,
			Message:  req.Message,
			Channel:  setting.Channel,
		})
	}

	return s.repo.InsertNotifications(notifs)
}
