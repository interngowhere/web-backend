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
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		res.Error = api.BuildError(err, WrapErrStrToInt, ListHandler)
		res.Message = WrapErrStrToInt.Message
		return res, err
	}

	c, err := GetCommentById(commentId)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveComments, UpdateHandler)
		res.Message = WrapErrRetrieveComments.Message
		return res, err
	}
	if c == nil {
		res.Error = api.BuildError(ErrNoMatchFromCommentID, WrapErrNoCommentFound, UpdateHandler)
		res.Message = WrapErrNoCommentFound.Message
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
func RemoveKudo(ctx context.Context, client *ent.Client, userId uuid.UUID, commentId int) (int, error) {
	return client.CommentKudo.
		Delete().
		Where(
			commentkudo.UserID(userId), 
			commentkudo.CommentID(commentId),
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

	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		res.Error = api.BuildError(err, WrapErrStrToInt, AddKudoHandler)
		res.Message = WrapErrStrToInt.Message
		return res, err
	}
	
	_, err = RemoveKudo(ctx, database.Client, userId, commentId)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRemoveKudo, DeleteHandler)
		res.Message = WrapErrRemoveKudo.Message
		return res, err
	}

	res.Message = SuccessfulRemoveKudoMessage

	return res, err
}