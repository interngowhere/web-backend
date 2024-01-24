package users

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
)

const DeleteHandler = "users.HandleDelete"
const SuccessfulDeleteUserMessage = "Successfully deleted user"

// DeleteUser deletes an existing user entry in the database
// based on ent.User object provided
func DeleteUser(ctx context.Context, client *ent.Client, userID uuid.UUID) error {
	return client.User.
		DeleteOneID(userID).
		Exec(ctx)
}

// HandleDelete handles the DELETE request, calls
// DeleteUser if needed and returns a JSON encoded API response
func HandleDelete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	userID, err := uuid.Parse(chi.URLParam(r, "userID"))
	if err != nil {
		res = api.BuildError(err, WrapErrParseUserID, DeleteHandler)
		return res, err
	}

	u, err := GetUserFromID(ctx, database.Client, userID)
	if err != nil {
		res = api.BuildError(err, WrapErrGetUser, DeleteHandler)
		return res, err
	}
	if u == nil {
		res = api.BuildError(customerrors.ErrResourceNotFound, customerrors.WrapErrNotFound, DeleteHandler)
		res.Message = customerrors.WrapErrNotFound.Message
		return res, customerrors.ErrResourceNotFound
	}

	err = DeleteUser(ctx, database.Client, userID)
	if err != nil {
		res = api.BuildError(err, WrapErrDeleteUser, DeleteHandler)
		return res, err
	}

	res.Message = SuccessfulDeleteUserMessage

	return res, err
}
