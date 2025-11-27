package service

import (
	"testing"

	"github.com/vidsrvcom/gotemp/internal/models"
	"github.com/vidsrvcom/gotemp/internal/repository"
)

func TestUserService_CreateUser(t *testing.T) {
	repo := repository.NewUserRepository()
	svc := NewUserService(repo)

	tests := []struct {
		name    string
		req     *models.CreateUserRequest
		wantErr bool
	}{
		{
			name: "valid request",
			req: &models.CreateUserRequest{
				Email: "test@example.com",
				Name:  "Test User",
			},
			wantErr: false,
		},
		{
			name: "missing email",
			req: &models.CreateUserRequest{
				Name: "Test User",
			},
			wantErr: true,
		},
		{
			name: "missing name",
			req: &models.CreateUserRequest{
				Email: "test@example.com",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := svc.CreateUser(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && user == nil {
				t.Error("CreateUser() user should not be nil")
			}
		})
	}
}

func TestUserService_GetUser(t *testing.T) {
	repo := repository.NewUserRepository()
	svc := NewUserService(repo)

	// Create a user first
	created, _ := svc.CreateUser(&models.CreateUserRequest{
		Email: "test@example.com",
		Name:  "Test User",
	})

	// Test getting existing user
	user, err := svc.GetUser(created.ID)
	if err != nil {
		t.Fatalf("GetUser() error = %v", err)
	}
	if user.ID != created.ID {
		t.Errorf("GetUser() ID = %v, want %v", user.ID, created.ID)
	}

	// Test getting non-existing user
	_, err = svc.GetUser(999)
	if err != models.ErrNotFound {
		t.Errorf("GetUser() error = %v, want %v", err, models.ErrNotFound)
	}
}

func TestUserService_GetAllUsers(t *testing.T) {
	repo := repository.NewUserRepository()
	svc := NewUserService(repo)

	// Create multiple users
	for i := 0; i < 3; i++ {
		_, _ = svc.CreateUser(&models.CreateUserRequest{
			Email: "test@example.com",
			Name:  "Test User",
		})
	}

	users, err := svc.GetAllUsers()
	if err != nil {
		t.Fatalf("GetAllUsers() error = %v", err)
	}

	if len(users) != 3 {
		t.Errorf("GetAllUsers() returned %d users, want 3", len(users))
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	repo := repository.NewUserRepository()
	svc := NewUserService(repo)

	created, _ := svc.CreateUser(&models.CreateUserRequest{
		Email: "test@example.com",
		Name:  "Test User",
	})

	updated, err := svc.UpdateUser(created.ID, &models.UpdateUserRequest{
		Name: "Updated Name",
	})
	if err != nil {
		t.Fatalf("UpdateUser() error = %v", err)
	}

	if updated.Name != "Updated Name" {
		t.Errorf("UpdateUser() Name = %v, want Updated Name", updated.Name)
	}

	// Test updating non-existing user
	_, err = svc.UpdateUser(999, &models.UpdateUserRequest{Name: "Test"})
	if err != models.ErrNotFound {
		t.Errorf("UpdateUser() error = %v, want %v", err, models.ErrNotFound)
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	repo := repository.NewUserRepository()
	svc := NewUserService(repo)

	created, _ := svc.CreateUser(&models.CreateUserRequest{
		Email: "test@example.com",
		Name:  "Test User",
	})

	err := svc.DeleteUser(created.ID)
	if err != nil {
		t.Fatalf("DeleteUser() error = %v", err)
	}

	_, err = svc.GetUser(created.ID)
	if err != models.ErrNotFound {
		t.Errorf("DeleteUser() user should not exist")
	}
}
