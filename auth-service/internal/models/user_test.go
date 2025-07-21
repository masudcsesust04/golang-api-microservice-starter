package models

import (
	"os"
	"testing"
	"time"

	"github.com/masudcsesust04/golang-jwt-auth/internal/config"
)

func TestMain(m *testing.M) {
	cleanup := config.TestDBCleaner(m)
	defer cleanup()

	os.Exit(m.Run())
}

func TestCreateAndGetUser(t *testing.T) {
	user := &User{
		FirstName:   "Test",
		LastName:    "User",
		PhoneNumber: "1234567890",
		Email:       "testuser@example.com",
		Status:      "active",
		Password:    "password123",
	}

	err := user.RegisterUser(user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	// Use GetUserByEmail instead of GetUserByUsername
	gotUser, err := user.GetUserByEmail(user.Email)
	if err != nil {
		t.Fatalf("GetUserByEmail failed: %v", err)
	}
	if gotUser == nil || gotUser.Email != user.Email {
		t.Fatalf("GetUserByEmail returned wrong user")
	}
}

func TestCreateAndDeleteRefreshToken(t *testing.T) {
	user := &User{
		FirstName:   "Token",
		LastName:    "User",
		PhoneNumber: "4445556666",
		Email:       "tokenuser@example.com",
		Password:    "password123",
	}

	err := user.RegisterUser(user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	rt := &RefreshToken{
		UserID:    user.ID,
		Token:     "testtoken",
		ExpiresAt: time.Now().Add(24 * time.Hour),
		CreatedAt: time.Now(),
	}

	err = user.CreateRefreshToken(rt)
	if err != nil {
		t.Fatalf("CreateRefreshToken failed: %v", err)
	}

	gotRT, err := user.GetRefreshToken(rt.UserID)
	if err != nil {
		t.Fatalf("GetRefreshToken failed: %v", err)
	}
	if gotRT == nil || gotRT.Token != rt.Token {
		t.Fatalf("GetRefreshToken returned wrong token")
	}

	err = user.DeleteRefreshToken(rt.UserID)
	if err != nil {
		t.Fatalf("DeleteRefreshToken failed: %v", err)
	}

	deletedRT, err := user.GetRefreshToken(rt.UserID)
	if err == nil && deletedRT != nil {
		t.Fatalf("DeleteRefreshToken did not delete token")
	}
}
