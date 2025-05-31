package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

func RandomUserAvatar(avatar string) string {
	return fmt.Sprintf("https://api.dicebear.com/6.x/initials/svg?seed=%s", avatar)
}

func GenerateOTP(length int) string {
	digits := "0123456789"
	var sb strings.Builder

	for range length {
		sb.WriteByte(digits[rand.Intn(len(digits))])
	}

	return sb.String()
}

func GenerateSlug(input string) string {

	slug := strings.ToLower(input)
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	suffix := strconv.Itoa(rand.Intn(1_000_000))
	slug = slug + "-" + leftPad(suffix, "0", 6)

	return slug
}

func leftPad(s string, pad string, length int) string {
	for len(s) < length {
		s = pad + s
	}
	return s
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateBase64QR(data string) string {
	var png []byte
	png, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(png)
}

const charset = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

func GenerateVerificationCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

type QRPayload struct {
	UserID     string `json:"userId"`
	BookingID  string `json:"bookingId"`
	ScheduleID string `json:"scheduleId"`
}

func ParseQRPayload(base64QR string) (*QRPayload, error) {
	decoded, err := base64.StdEncoding.DecodeString(base64QR)
	if err != nil {
		return nil, err
	}

	var payload QRPayload
	if err := json.Unmarshal(decoded, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func GenerateInvoiceNumber(paymentID uuid.UUID) string {
	timestamp := time.Now().Format("20060102")
	shortID := paymentID.String()[:8]
	return fmt.Sprintf("INV/%s/%s", timestamp, shortID)
}
