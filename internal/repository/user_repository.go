// Package repository provides data access implementations.
package repository

import (
	"sync"
	"time"

	"github.com/vidsrvcom/gotemp/internal/models"
)

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id int64) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(user *models.User) error
	Delete(id int64) error
}

// userRepository implements UserRepository with in-memory storage.
type userRepository struct {
	mu     sync.RWMutex
	users  map[int64]*models.User
	nextID int64
}

// NewUserRepository creates a new user repository instance.
func NewUserRepository() UserRepository {
	return &userRepository{
		users:  make(map[int64]*models.User),
		nextID: 1,
	}
}

// Create adds a new user to the repository.
func (r *userRepository) Create(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	r.users[user.ID] = user
	r.nextID++

	return nil
}

// GetByID retrieves a user by their ID.
func (r *userRepository) GetByID(id int64) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, models.ErrNotFound
	}

	return user, nil
}

// GetAll retrieves all users from the repository.
func (r *userRepository) GetAll() ([]*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

// Update modifies an existing user in the repository.
func (r *userRepository) Update(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return models.ErrNotFound
	}

	user.UpdatedAt = time.Now()
	r.users[user.ID] = user

	return nil
}

// Delete removes a user from the repository.
func (r *userRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return models.ErrNotFound
	}

	delete(r.users, id)
	return nil
}
