package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/vidsrvcom/gotemp/internal/models"
	"github.com/vidsrvcom/gotemp/internal/repository"
	"github.com/vidsrvcom/gotemp/internal/service"
)

func setupTestHandler() (*UserHandler, *chi.Mux) {
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)
	handler := NewUserHandler(svc)

	r := chi.NewRouter()
	r.Post("/api/users", handler.Create)
	r.Get("/api/users", handler.GetAll)
	r.Get("/api/users/{id}", handler.Get)
	r.Put("/api/users/{id}", handler.Update)
	r.Delete("/api/users/{id}", handler.Delete)

	return handler, r
}

func TestUserHandler_Create(t *testing.T) {
	_, router := setupTestHandler()

	tests := []struct {
		name       string
		body       interface{}
		wantStatus int
	}{
		{
			name: "valid request",
			body: models.CreateUserRequest{
				Email: "test@example.com",
				Name:  "Test User",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "missing email",
			body: models.CreateUserRequest{
				Name: "Test User",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid json",
			body:       "invalid",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("Create() status = %v, want %v", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestUserHandler_Get(t *testing.T) {
	_, router := setupTestHandler()

	// Create a user first
	createBody, _ := json.Marshal(models.CreateUserRequest{
		Email: "test@example.com",
		Name:  "Test User",
	})
	createReq := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(createBody))
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()
	router.ServeHTTP(createRec, createReq)

	var created models.User
	json.Unmarshal(createRec.Body.Bytes(), &created)

	tests := []struct {
		name       string
		id         string
		wantStatus int
	}{
		{
			name:       "existing user",
			id:         "1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "non-existing user",
			id:         "999",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "invalid id",
			id:         "invalid",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/users/"+tt.id, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("Get() status = %v, want %v", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestUserHandler_GetAll(t *testing.T) {
	_, router := setupTestHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("GetAll() status = %v, want %v", rec.Code, http.StatusOK)
	}
}

func TestUserHandler_Update(t *testing.T) {
	_, router := setupTestHandler()

	// Create a user first
	createBody, _ := json.Marshal(models.CreateUserRequest{
		Email: "test@example.com",
		Name:  "Test User",
	})
	createReq := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(createBody))
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()
	router.ServeHTTP(createRec, createReq)

	updateBody, _ := json.Marshal(models.UpdateUserRequest{
		Name: "Updated Name",
	})

	req := httptest.NewRequest(http.MethodPut, "/api/users/1", bytes.NewReader(updateBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Update() status = %v, want %v", rec.Code, http.StatusOK)
	}
}

func TestUserHandler_Delete(t *testing.T) {
	_, router := setupTestHandler()

	// Create a user first
	createBody, _ := json.Marshal(models.CreateUserRequest{
		Email: "test@example.com",
		Name:  "Test User",
	})
	createReq := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(createBody))
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()
	router.ServeHTTP(createRec, createReq)

	req := httptest.NewRequest(http.MethodDelete, "/api/users/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("Delete() status = %v, want %v", rec.Code, http.StatusNoContent)
	}
}

func TestHomeHandler_Health(t *testing.T) {
	handler := NewHomeHandler()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler.Health(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Health() status = %v, want %v", rec.Code, http.StatusOK)
	}

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)

	if response["status"] != "healthy" {
		t.Errorf("Health() status = %v, want healthy", response["status"])
	}
}
