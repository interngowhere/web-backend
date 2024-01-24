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
// a valid JWT and returns the parsed user id or
// a nil UUID if an error occurs.
func GetUserIDFromToken(r *http.Request) (uuid.UUID, error) {
	token, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return uuid.Nil, err
	}

	id, _ := token.Get("id")
	userID, err := uuid.Parse(fmt.Sprintf("%s", id))
	if err != nil {
		return uuid.Nil, err
	}
	return userID, err
}

// GetUserFromID queries the database for a user with the given ID
func GetUserFromID(ctx context.Context, client *ent.Client, id uuid.UUID) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}
