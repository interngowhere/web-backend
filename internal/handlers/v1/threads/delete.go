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
	threadId, err := strconv.Atoi(chi.URLParam(r, "threadId"))
	if err != nil {
		res.Error = api.BuildError(err, WrapErrStrToInt, UpdateHandler)
		res.Message = WrapErrStrToInt.Message
		return res, err
	}
	t, err := GetThreads(threadId)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveThreads, DeleteHandler)
		res.Message = WrapErrRetrieveThreads.Message
		return res, err
	}
	if len(t) == 0 {
		res.Error = api.BuildError(ErrNoMatchFromID, WrapErrNoThreadFound, DeleteHandler)
		res.Message = WrapErrNoThreadFound.Message
		return res, ErrNoMatchFromID
	}

	err = DeleteThread(ctx, database.Client, t[0])
	if err != nil {
		res.Error = api.BuildError(err, WrapErrDeleteThread, DeleteHandler)
		res.Message = WrapErrDeleteThread.Message
		return res, err
	}

	res.Message = SuccessfulDeleteThreadMessage

	return res, err
}

// RemoveKudo deletes an existing kudo entry consisting user_id and thread_id
func RemoveKudo(ctx context.Context, client *ent.Client, userId uuid.UUID, threadId int) (int, error) {
	return client.ThreadKudo.
		Delete().
		Where(
			threadkudo.UserID(userId),
			threadkudo.ThreadID(threadId),
		).
		Exec(ctx)
}

// HandleRemoveKudo handles the DELETE request, calls
// RemoveKudo if there is no error and returns a JSON encoded API response
func HandleRemoveKudo(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	userId, err := users.GetUserIDFromToken(r)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveUserID, CreateHandler)
		res.Message = WrapErrRetrieveUserID.Message
		return res, err
	}

	threadId, err := strconv.Atoi(chi.URLParam(r, "threadId"))
	if err != nil {
		res.Error = api.BuildError(err, WrapErrStrToInt, AddKudoHandler)
		res.Message = WrapErrStrToInt.Message
		return res, err
	}

	_, err = RemoveKudo(ctx, database.Client, userId, threadId)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRemoveKudo, DeleteHandler)
		res.Message = WrapErrRemoveKudo.Message
		return res, err
	}

	res.Message = SuccessfulRemoveKudoMessage

	return res, err
}
