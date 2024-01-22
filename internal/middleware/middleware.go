package middleware

import (
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{os.Getenv("PROD_ORIGIN"), os.Getenv("DEV_ORIGIN")},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))
}
