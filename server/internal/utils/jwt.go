package utils

import (
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var accessTokenSecret = []byte(os.Getenv("JWT_ACCESS_SECRET"))
var refreshTokenSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))

type Claims struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessTokenSecret)
}

func GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshTokenSecret)
}

func DecodeAccessToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return accessTokenSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid access token")
}

func DecodeRefreshToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenSecret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims.Subject, nil
	}
	return "", errors.New("invalid refresh token")
}

func SetRefreshTokenCookie(c *gin.Context, refreshToken string) {
	c.SetCookie(
		"refreshToken",
		refreshToken,
		7*24*3600, // 7 hari
		"/",
		"",   // domain kosong (pakai domain API)
		true, // Secure
		true, // HttpOnly
	)
}

func ClearRefreshTokenCookie(c *gin.Context) {
	c.SetCookie(
		"refreshToken",
		"",
		-1, // expire segera
		"/",
		"",
		true,
		true,
	)
}

func SetAccessTokenCookie(c *gin.Context, accessToken string) {
	c.SetCookie(
		"accessToken",
		accessToken,
		1*3600, // 1 jam
		"/",
		"",
		true,
		true,
	)
}

func ClearAccessTokenCookie(c *gin.Context) {
	c.SetCookie(
		"accessToken",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)
}
