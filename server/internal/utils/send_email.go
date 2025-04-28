package utils

import (
	"fmt"
	"os"

	"server/internal/config"

	gomail "gopkg.in/gomail.v2"
)

func SendEmail(email, otpcode, body string) error {
	m := gomail.NewMessage()

	from := os.Getenv("USER_EMAIL")
	subject := "One-Time Password (OTP) xxxxx"

	m.SetHeader("From", fmt.Sprintf("fitness_app <%s>", from))
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	m.AddAlternative("text/html", fmt.Sprintf("Your OTP Code is <b>%s</b>", otpcode))

	if err := config.MailDialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
