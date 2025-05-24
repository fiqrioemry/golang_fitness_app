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
	attendanceService   services.AttendanceService
}

func NewCronManager(
	payment services.PaymentService,
	schedule services.ScheduleTemplateService,
	notification services.NotificationService,
	attendance services.AttendanceService,
) *CronManager {
	return &CronManager{
		c:                   cron.New(cron.WithSeconds()),
		paymentService:      payment,
		scheduleService:     schedule,
		notificationService: notification,
		attendanceService:   attendance,
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
	cm.c.AddFunc("0 12 * * *", func() {
		log.Println("Cron: Auto-generating class schedules...")
		if err := cm.scheduleService.AutoGenerateSchedules(); err != nil {
			log.Println("Schedule generation failed:", err)
		} else {
			log.Println("Recurring class schedules generated")
		}
	})

	// Reminder dan Mark Absent  tiap 15 menit
	cm.c.AddFunc("@every 15m", func() {
		log.Println("Cron: Sending class reminders...")
		if err := cm.notificationService.SendClassReminder(); err != nil {
			log.Println(" Reminder failed:", err)
		} else {
			log.Println("Class reminders sent")
		}

		log.Println("Cron: Marking absents for missed bookings...")
		if err := cm.attendanceService.MarkAbsentBookings(); err != nil {
			log.Println(" Marking absents failed:", err)
		} else {
			log.Println("Marked absent attendees")
		}
	})

}

func (cm *CronManager) Start() {
	cm.c.Start()
	log.Println("Cron Manager started")
}
