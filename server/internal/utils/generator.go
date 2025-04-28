package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func RandomUserAvatar() string {
	avatars := []string{
		"Shopia", "Jameson", "Emery", "Sawyer", "Maria",
		"Chase", "Jocelyn", "Liliana", "Robert", "Christian", "Nolan",
	}

	rand.Seed(time.Now().UnixNano())
	avatar := avatars[rand.Intn(len(avatars))]

	return fmt.Sprintf("https://api.dicebear.com/9.x/fun-emoji/svg?seed=%s", avatar)
}

func GenerateOTP(length int) string {
	digits := "0123456789"
	var sb strings.Builder

	for i := 0; i < length; i++ {
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
