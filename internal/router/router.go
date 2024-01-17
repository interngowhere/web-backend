package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/internal/routes"
)

// Setup mounts all handlers
func Setup(r chi.Router) {
	// Public Routes
	r.Group(routes.PublicRoutes())

	// Protected Routes
	r.Group(routes.ProtectedRoutes())
}
