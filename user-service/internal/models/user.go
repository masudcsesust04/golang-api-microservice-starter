package models

import (
	"context"
	"fmt"
	"time"

	"github.com/masudcsesust04/golang-jwt-auth/internal/config"
)

// User represent a user in the system
type User struct {
	ID          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GetUserByID retrives a user by ID
func (u *User) GetUserByID(id int64) (*User, error) {
	query := `SELECT id, first_name, last_name, phone_number, email, status, created_at, updated_at FROM  users WHERE id = $1`
	user := &User{}

	err := config.DbConn.GetPool().QueryRow(context.Background(), query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Email, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return user, nil
}

// GetAllUsers retrives all users from the database
func (u *User) GetAllUsers() ([]*User, error) {
	query := `SELECT id, first_name, last_name, phone_number, email, status, created_at, updated_at FROM  users`

	rows, err := config.DbConn.GetPool().Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Email, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}

// UpdateUser updates an existing users' information
func (u *User) UpdateUser(user *User) error {
	query := `UPDATE users SET first_name = $1, last_name = $2, phone_number = $3, email = $4, status= $5 WHERE id = $6`
	_, err := config.DbConn.GetPool().Exec(context.Background(), query, user.FirstName, user.LastName, user.PhoneNumber, user.Email, user.Status, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// DeleteUser deletes a user by ID
func (u *User) DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE ID = $1`
	_, err := config.DbConn.GetPool().Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
