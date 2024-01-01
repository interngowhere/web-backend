package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/internal/routes"
)

// setUpRoutes mounts all handlers
func Setup(r chi.Router) {
	// Public Routes
	r.Group(routes.GetRoutes())

	// TODO Private Routes
	// r.Use(AuthMiddleware)
    // r.Group(routes.PostRoutes())
}
