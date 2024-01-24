package threads

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/threadkudo"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
	"github.com/interngowhere/web-backend/internal/handlers/v1/users"
)

const DeleteHandler = "threads.HandleDelete"
const RemoveKudoHandler = "threads.HandleRemoveKudo"
const SuccessfulDeleteThreadMessage = "Successfully deleted thread"
const SuccessfulRemoveKudoMessage = "Successfully removed kudo on the given thread"

// DeleteThread deletes an existing thread entry in the database
// based on ent.Thread object provided
func DeleteThread(ctx context.Context, client *ent.Client, thread *ent.Thread) error {
	return client.Thread.
		DeleteOne(thread).
		Exec(ctx)
}

// HandleDelete handles the DELETE request, calls
// DeleteThread if needed and returns a JSON encoded API response
func HandleDelete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	// Retrieve a reference to the thread
	threadID, err := strconv.Atoi(chi.URLParam(r, "threadID"))
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrStrToInt, UpdateHandler)
		return res, err
	}
	t, err := GetThreadByID(ctx, threadID)
	if err != nil {
		res = api.BuildError(err, WrapErrRetrieveThreads, DeleteHandler)
		return res, err
	}
	if t == nil {
		res = api.BuildError(customerrors.ErrResourceNotFound, customerrors.WrapErrNotFound, DeleteHandler)
		res.Message = customerrors.WrapErrNotFound.Message
		return res, customerrors.ErrResourceNotFound
	}

	err = DeleteThread(ctx, database.Client, t)
	if err != nil {
		res = api.BuildError(err, WrapErrDeleteThread, DeleteHandler)
		return res, err
	}

	res.Message = SuccessfulDeleteThreadMessage

	return res, err
}

// RemoveKudo deletes an existing kudo entry consisting user_id and thread_id
func RemoveKudo(ctx context.Context, client *ent.Client, userID uuid.UUID, threadID int) (int, error) {
	return client.ThreadKudo.
		Delete().
		Where(
			threadkudo.UserID(userID),
			threadkudo.ThreadID(threadID),
		).
		Exec(ctx)
}

// HandleRemoveKudo handles the DELETE request, calls
// RemoveKudo if there is no error and returns a JSON encoded API response
func HandleRemoveKudo(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	userID, err := users.GetUserIDFromToken(r)
	if err != nil {
		res = api.BuildError(err, users.WrapErrRetrieveIDFromJWT, CreateHandler)
		return res, err
	}

	threadID, err := strconv.Atoi(chi.URLParam(r, "threadID"))
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrStrToInt, AddKudoHandler)
		return res, err
	}

	_, err = RemoveKudo(ctx, database.Client, userID, threadID)
	if err != nil {
		res = api.BuildError(err, WrapErrRemoveKudo, DeleteHandler)
		return res, err
	}

	res.Message = SuccessfulRemoveKudoMessage

	return res, err
}
