// Package models defines the data structures used throughout the application.
package models

import (
	"time"
)

// User represents a user in the system.
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request payload for creating a user.
type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Validate validates the CreateUserRequest fields.
func (r *CreateUserRequest) Validate() error {
	if r.Email == "" {
		return ErrEmailRequired
	}
	if r.Name == "" {
		return ErrNameRequired
	}
	return nil
}

// UpdateUserRequest represents the request payload for updating a user.
type UpdateUserRequest struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}
