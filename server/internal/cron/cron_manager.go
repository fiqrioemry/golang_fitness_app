package cron

import (
	"log"

	"server/internal/services"

	"github.com/robfig/cron/v3"
)

type CronManager struct {
	c                   *cron.Cron
	paymentService      services.PaymentService
	scheduleService     services.ScheduleTemplateService
	notificationService services.NotificationService
}

func NewCronManager(
	payment services.PaymentService,
	schedule services.ScheduleTemplateService,
	notification services.NotificationService,
) *CronManager {
	return &CronManager{
		c:                   cron.New(cron.WithSeconds()),
		paymentService:      payment,
		scheduleService:     schedule,
		notificationService: notification,
	}
}

func (cm *CronManager) RegisterJobs() {
	// Update payment pending → failed (setiap hari 00:30)
	cm.c.AddFunc("0 30 0 * * *", func() {
		log.Println("Cron: Checking expired pending payments...")
		if err := cm.paymentService.ExpireOldPendingPayments(); err != nil {
			log.Println("Error expiring payments:", err)
		} else {
			log.Println("Payment status updated (pending → failed)")
		}
	})

	// Generate recurring class schedule (setiap hari 12:00)
	cm.c.AddFunc("0 0 12 * * *", func() {
		log.Println("Cron: Auto-generating class schedules...")
		if err := cm.scheduleService.AutoGenerateSchedules(); err != nil {
			log.Println("Schedule generation failed:", err)
		} else {
			log.Println("Recurring class schedules generated")
		}
	})

	// Kirim notifikasi reminder kelas (setiap 15 menit)
	cm.c.AddFunc("@every 15m", func() {
		log.Println("Cron: Sending class reminders...")
		if err := cm.notificationService.SendClassReminder(); err != nil {
			log.Println("Failed to send class reminders:", err)
		} else {
			log.Println("Class reminders sent successfully")
		}
	})
}

func (cm *CronManager) Start() {
	cm.c.Start()
	log.Println("Cron Manager started")
}
