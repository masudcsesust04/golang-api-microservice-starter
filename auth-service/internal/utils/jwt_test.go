package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateTestToken(t *testing.T, secret string) string {
	t.Helper()
	claims := jwt.MapClaims{
		"user_id": 1,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Failed to generate test token: %v", err)
	}

	return tokenString
}
