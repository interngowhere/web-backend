package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/user"
)

// GetUserFromEmail queries the database for a user with
// the provided email address and returns a user object
func GetUserFromEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error) {
	return client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
}

// GetUserIDFromToken checks request context for
// a valid JWT and returns the parsed user id
func GetUserIDFromToken(r *http.Request) (uuid.UUID, error) {
	token, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return uuid.Nil, err
	}

	id, _ := token.Get("id")
	userId, err := uuid.Parse(fmt.Sprintf("%s", id))
	if err != nil {
		return uuid.Nil, err
	}
	return userId, err
}
