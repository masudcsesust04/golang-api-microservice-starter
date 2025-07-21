package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/masudcsesust04/golang-jwt-auth/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey string

func GenerateAccessToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}

func GenerateRefreshToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		// fallback to less secure random string if needed
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

func ValidateRefreshToken(refreshToken *models.RefreshToken) error {
	if time.Now().After(refreshToken.ExpiresAt) {
		return fmt.Errorf("refresh token has expired")
	}

	return nil
}

func GenerateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

func HashToken(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareToken(hash, token string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
}

func SetJWTSecrectKey(secret string) {
	jwtSecretKey = secret
}
