// Package server provides HTTP server configuration and routing.
package server

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/vidsrvcom/gotemp/internal/config"
	"github.com/vidsrvcom/gotemp/internal/handler"
	"github.com/vidsrvcom/gotemp/internal/middleware"
)

// Server represents the HTTP server.
type Server struct {
	cfg         *config.Config
	logger      *log.Logger
	httpServer  *http.Server
	homeHandler *handler.HomeHandler
	userHandler *handler.UserHandler
}

// New creates a new Server instance.
func New(cfg *config.Config, logger *log.Logger, homeHandler *handler.HomeHandler, userHandler *handler.UserHandler) *Server {
	return &Server{
		cfg:         cfg,
		logger:      logger,
		homeHandler: homeHandler,
		userHandler: userHandler,
	}
}

// Start begins listening for HTTP requests.
func (s *Server) Start() error {
	router := s.setupRoutes()

	s.httpServer = &http.Server{
		Addr:    s.cfg.ServerAddress(),
		Handler: router,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// setupRoutes configures all application routes.
func (s *Server) setupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(middleware.Logger(s.logger))
	r.Use(middleware.Recoverer(s.logger))
	r.Use(middleware.CORS)

	// Static files
	fileServer := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Home routes
	r.Get("/", s.homeHandler.Index)
	r.Get("/health", s.homeHandler.Health)

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.ContentType("application/json"))

		r.Route("/users", func(r chi.Router) {
			r.Get("/", s.userHandler.GetAll)
			r.Post("/", s.userHandler.Create)
			r.Get("/{id}", s.userHandler.Get)
			r.Put("/{id}", s.userHandler.Update)
			r.Delete("/{id}", s.userHandler.Delete)
		})
	})

	return r
}
