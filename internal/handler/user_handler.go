// Package handler contains HTTP request handlers.
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vidsrvcom/gotemp/internal/models"
	"github.com/vidsrvcom/gotemp/internal/service"
)

// UserHandler handles HTTP requests for user operations.
type UserHandler struct {
	service service.UserService
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{service: svc}
}

// Create handles POST /api/users - creates a new user.
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := h.service.CreateUser(&req)
	if err != nil {
		switch err {
		case models.ErrEmailRequired, models.ErrNameRequired:
			respondError(w, http.StatusBadRequest, err.Error())
		default:
			respondError(w, http.StatusInternalServerError, "Failed to create user")
		}
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

// Get handles GET /api/users/{id} - retrieves a user by ID.
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		if err == models.ErrNotFound {
			respondError(w, http.StatusNotFound, "User not found")
			return
		}
		respondError(w, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}

	respondJSON(w, http.StatusOK, user)
}

// GetAll handles GET /api/users - retrieves all users.
func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	respondJSON(w, http.StatusOK, users)
}

// Update handles PUT /api/users/{id} - updates an existing user.
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := h.service.UpdateUser(id, &req)
	if err != nil {
		if err == models.ErrNotFound {
			respondError(w, http.StatusNotFound, "User not found")
			return
		}
		respondError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	respondJSON(w, http.StatusOK, user)
}

// Delete handles DELETE /api/users/{id} - deletes a user.
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		if err == models.ErrNotFound {
			respondError(w, http.StatusNotFound, "User not found")
			return
		}
		respondError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
