package comments

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/commentkudo"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
	"github.com/interngowhere/web-backend/internal/handlers/v1/users"
)

const DeleteHandler = "comments.HandleDelete"
const RemoveKudoHandler = "comments.HandleRemoveKudo"
const SuccessfulDeleteCommentMessage = "Successfully deleted comment"
const SuccessfulRemoveKudoMessage = "Successfully removed kudo from comment"

// DeleteComment deletes an existing comment entry in the database
// based on ent.Comment object provided
func DeleteComment(ctx context.Context, client *ent.Client, comment *ent.Comment) error {
	return client.Comment.
		DeleteOne(comment).
		Exec(ctx)
}

// HandleDelete handles the DELETE request, calls
// DeleteComment if needed and returns a JSON encoded API response
func HandleDelete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	// Retrieve a reference to the comment
	commentID, err := strconv.Atoi(chi.URLParam(r, "commentID"))
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrStrToInt, ListHandler)
		res.Message = customerrors.WrapErrStrToInt.Message
		return res, err
	}

	c, err := GetCommentById(ctx, commentID)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveComments, UpdateHandler)
		res.Message = WrapErrRetrieveComments.Message
		return res, err
	}
	if c == nil {
		res.Error = api.BuildError(ErrNoMatchFromCommentID, customerrors.WrapErrNotFound, UpdateHandler)
		res.Message = customerrors.WrapErrNotFound.Message
		return res, ErrNoMatchFromCommentID
	}

	err = DeleteComment(ctx, database.Client, c)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrDeleteComment, DeleteHandler)
		res.Message = WrapErrDeleteComment.Message
		return res, err
	}

	res.Message = SuccessfulDeleteCommentMessage

	return res, err
}

// RemoveKudo deletes an existing kudo entry consisting user_id and comment_id
func RemoveKudo(ctx context.Context, client *ent.Client, userID uuid.UUID, commentID int) (int, error) {
	return client.CommentKudo.
		Delete().
		Where(
			commentkudo.UserID(userID),
			commentkudo.CommentID(commentID),
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
		res.Error = api.BuildError(err, users.WrapErrRetrieveIDFromJWT, CreateHandler)
		res.Message = users.WrapErrRetrieveIDFromJWT.Message
		return res, err
	}

	commentID, err := strconv.Atoi(chi.URLParam(r, "commentID"))
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrStrToInt, AddKudoHandler)
		res.Message = customerrors.WrapErrStrToInt.Message
		return res, err
	}

	_, err = RemoveKudo(ctx, database.Client, userID, commentID)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRemoveKudo, DeleteHandler)
		res.Message = WrapErrRemoveKudo.Message
		return res, err
	}

	res.Message = SuccessfulRemoveKudoMessage

	return res, err
}
