package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/interngowhere/web-backend/internal/api"
	jwt "github.com/interngowhere/web-backend/internal/auth"
	"github.com/interngowhere/web-backend/internal/handlers/v1/comments"
	"github.com/interngowhere/web-backend/internal/handlers/v1/tags"
	"github.com/interngowhere/web-backend/internal/handlers/v1/threads"
	"github.com/interngowhere/web-backend/internal/handlers/v1/topics"
	"github.com/interngowhere/web-backend/internal/handlers/v1/users"
)

func PublicRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		// Topics
		r.Get("/topics", api.BuildRouteHandler(topics.HandleList))
		r.Get("/topics/{title}", api.BuildRouteHandler(topics.HandleList))

		// Threads
		r.Get("/threads", api.BuildRouteHandler(threads.HandleList))
		r.Get("/threads/{threadID}", api.BuildRouteHandler(threads.HandleList))

		// Comments
		r.Get("/threads/{threadID}/comments", api.BuildRouteHandler(comments.HandleList))

		// Tags
		r.Get("/tags", api.BuildRouteHandler(tags.HandleList))

		// Users
		r.Post("/users", api.BuildRouteHandler(users.HandleCreate))
		r.Post("/login", api.BuildRouteHandler(users.HandleLogin))
	}
}

func ProtectedRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(jwtauth.Authenticator(jwt.TokenAuth))

		// Check if user is still authenticated
		r.Post("/ping", func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("."))
		})

		// Users
		r.Put("/users/{userID}", api.BuildRouteHandler(users.HandleUpdate))
		r.Delete("/users/{userID}", api.BuildRouteHandler(users.HandleDelete))

		// Topics
		r.Post("/topics", api.BuildRouteHandler(topics.HandleCreate))
		r.Put("/topics/{title}", api.BuildRouteHandler(topics.HandleUpdate))
		r.Delete("/topics/{title}", api.BuildRouteHandler(topics.HandleDelete))

		// Threads
		r.Post("/topics/{title}/threads", api.BuildRouteHandler(threads.HandleCreate))
		r.Put("/threads/{threadID}", api.BuildRouteHandler(threads.HandleUpdate))
		r.Delete("/threads/{threadID}", api.BuildRouteHandler(threads.HandleDelete))
		r.Post("/threads/{threadID}/kudo", api.BuildRouteHandler(threads.HandleAddKudo))
		r.Delete("/threads/{threadID}/kudo", api.BuildRouteHandler(threads.HandleRemoveKudo))

		// Comments
		r.Post("/threads/{threadID}/comments", api.BuildRouteHandler(comments.HandleCreate))
		r.Put("/threads/{threadID}/comments/{commentID}", api.BuildRouteHandler(comments.HandleUpdate))
		r.Delete("/threads/{threadID}/comments/{commentID}", api.BuildRouteHandler(comments.HandleDelete))
		r.Post("/threads/{threadID}/comments/{commentID}/kudo", api.BuildRouteHandler(comments.HandleAddKudo))
		r.Delete("/threads/{threadID}/comments/{commentID}/kudo", api.BuildRouteHandler(comments.HandleRemoveKudo))

		// Tags
		r.Post("/tags", api.BuildRouteHandler(tags.HandleCreate))
		r.Delete("/tags/{tagId}", api.BuildRouteHandler(tags.HandleDelete))
	}
}
