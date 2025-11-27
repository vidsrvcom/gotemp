// Package service contains business logic implementations.
package service

import (
	"github.com/vidsrvcom/gotemp/internal/models"
	"github.com/vidsrvcom/gotemp/internal/repository"
)

// UserService defines the interface for user business operations.
type UserService interface {
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	GetUser(id int64) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(id int64, req *models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id int64) error
}

// userService implements UserService.
type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new user service instance.
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser creates a new user.
func (s *userService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user := &models.User{
		Email: req.Email,
		Name:  req.Name,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID.
func (s *userService) GetUser(id int64) (*models.User, error) {
	return s.repo.GetByID(id)
}

// GetAllUsers retrieves all users.
func (s *userService) GetAllUsers() ([]*models.User, error) {
	return s.repo.GetAll()
}

// UpdateUser updates an existing user.
func (s *userService) UpdateUser(id int64, req *models.UpdateUserRequest) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Name != "" {
		user.Name = req.Name
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID.
func (s *userService) DeleteUser(id int64) error {
	return s.repo.Delete(id)
}
