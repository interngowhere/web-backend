package middleware

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-chi/jwtauth/v5"
	jwt "github.com/interngowhere/web-backend/internal/auth"
)

// Setup mounts all middleware
func Setup(r chi.Router) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(httprate.LimitByIP(120, time.Minute))
	r.Use(jwtauth.Verifier(jwt.TokenAuth))
}
