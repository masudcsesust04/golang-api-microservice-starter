package models

import (
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	user := &User{}
	users, err := user.GetAllUsers()
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}
	if len(users) == 0 {
		t.Fatalf("GetAllUsers returned empty list")
	}
}
