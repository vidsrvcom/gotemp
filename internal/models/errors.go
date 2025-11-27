package models

import "errors"

// Common errors used across the application.
var (
	ErrNotFound       = errors.New("resource not found")
	ErrEmailRequired  = errors.New("email is required")
	ErrNameRequired   = errors.New("name is required")
	ErrInvalidRequest = errors.New("invalid request")
	ErrAlreadyExists  = errors.New("resource already exists")
)
