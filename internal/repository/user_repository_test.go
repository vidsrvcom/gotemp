package repository

import (
	"testing"

	"github.com/vidsrvcom/gotemp/internal/models"
)

func TestUserRepository_Create(t *testing.T) {
	repo := NewUserRepository()

	user := &models.User{
		Email: "test@example.com",
		Name:  "Test User",
	}

	err := repo.Create(user)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	if user.ID != 1 {
		t.Errorf("Create() ID = %v, want 1", user.ID)
	}

	if user.CreatedAt.IsZero() {
		t.Error("Create() CreatedAt should be set")
	}
}

func TestUserRepository_GetByID(t *testing.T) {
	repo := NewUserRepository()

	// Create a user first
	user := &models.User{
		Email: "test@example.com",
		Name:  "Test User",
	}
	_ = repo.Create(user)

	// Test getting existing user
	found, err := repo.GetByID(user.ID)
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}
	if found.Email != user.Email {
		t.Errorf("GetByID() Email = %v, want %v", found.Email, user.Email)
	}

	// Test getting non-existing user
	_, err = repo.GetByID(999)
	if err != models.ErrNotFound {
		t.Errorf("GetByID() error = %v, want %v", err, models.ErrNotFound)
	}
}

func TestUserRepository_GetAll(t *testing.T) {
	repo := NewUserRepository()

	// Create multiple users
	for i := 0; i < 3; i++ {
		user := &models.User{
			Email: "test@example.com",
			Name:  "Test User",
		}
		_ = repo.Create(user)
	}

	users, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() error = %v", err)
	}

	if len(users) != 3 {
		t.Errorf("GetAll() returned %d users, want 3", len(users))
	}
}

func TestUserRepository_Update(t *testing.T) {
	repo := NewUserRepository()

	user := &models.User{
		Email: "test@example.com",
		Name:  "Test User",
	}
	_ = repo.Create(user)

	user.Name = "Updated Name"
	err := repo.Update(user)
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	found, _ := repo.GetByID(user.ID)
	if found.Name != "Updated Name" {
		t.Errorf("Update() Name = %v, want Updated Name", found.Name)
	}

	// Test updating non-existing user
	nonExisting := &models.User{ID: 999}
	err = repo.Update(nonExisting)
	if err != models.ErrNotFound {
		t.Errorf("Update() error = %v, want %v", err, models.ErrNotFound)
	}
}

func TestUserRepository_Delete(t *testing.T) {
	repo := NewUserRepository()

	user := &models.User{
		Email: "test@example.com",
		Name:  "Test User",
	}
	_ = repo.Create(user)

	err := repo.Delete(user.ID)
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	_, err = repo.GetByID(user.ID)
	if err != models.ErrNotFound {
		t.Errorf("Delete() user should not exist, got error = %v", err)
	}

	// Test deleting non-existing user
	err = repo.Delete(999)
	if err != models.ErrNotFound {
		t.Errorf("Delete() error = %v, want %v", err, models.ErrNotFound)
	}
}
