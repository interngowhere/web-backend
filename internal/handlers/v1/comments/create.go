package comments

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/commentkudo"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
	"github.com/interngowhere/web-backend/internal/handlers/v1/threads"
	"github.com/interngowhere/web-backend/internal/handlers/v1/users"
)

const CreateHandler = "comments.HandleCreate"
const AddKudoHandler = "comments.HandleAddKudo"
const SuccessfulCreateCommentMessage = "Successfully created comment"
const SuccessfulAddKudoMessage = "Successfully added kudo to comment"

// CreateComment adds a new comment entry in the database
func CreateComment(ctx context.Context, client *ent.Client, threadID int, comment ent.Comment) error {
	return client.Comment.
		Create().
		SetParentID(comment.ParentID).
		SetContent(comment.Content).
		SetCreatedBy(comment.CreatedBy).
		SetCreatedAt(comment.CreatedAt).
		SetCommentAuthorsID(comment.CreatedBy).
		SetThreadsID(threadID).
		Exec(ctx)
}

// HandleCreate parses the POST request form data, calls
// CreateComment if needed and returns a JSON encoded API response
func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	userID, err := users.GetUserIDFromToken(r)
	if err != nil {
		res.Error = api.BuildError(err, users.WrapErrRetrieveIDFromJWT, CreateHandler)
		res.Message = users.WrapErrRetrieveIDFromJWT.Message
		return res, err
	}

	threadID, err := strconv.Atoi(chi.URLParam(r, "threadID"))
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrStrToInt, ListHandler)
		res.Message = customerrors.WrapErrStrToInt.Message
		return res, err
	}

	_, err = threads.GetThreadByID(ctx, threadID)
	if err != nil {
		res.Error = api.BuildError(err, threads.WrapErrRetrieveThreads, CreateHandler)
		res.Message = threads.WrapErrRetrieveThreads.Message
		return res, err
	}

	// Read JSON body in request into a new ThreadRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data CommentRequest
	err = decoder.Decode(&data)
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
		res.Message = customerrors.WrapErrDecodeRequest.Message
		return res, err
	}

	// Check if content field is present in request body
	// Note: content is the only required field in the POST request
	// If parentId field is not provided, parentId will be set to 0
	// by default, i.e., the comment has no parent.
	if len(data.Content) == 0 {
		res.Error = api.BuildError(customerrors.ErrMalformedRequest, customerrors.WrapErrRequestFormat, CreateHandler)
		res.Message = customerrors.WrapErrRequestFormat.Message
		return res, customerrors.ErrMalformedRequest
	}

	if len(data.ParentID) == 0 {
		data.ParentID = "0"
	}
	parentId, err := strconv.Atoi(data.ParentID)
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrStrToInt, CreateHandler)
		res.Message = customerrors.WrapErrStrToInt.Message
		return res, err
	}

	c := ent.Comment{
		ParentID:  parentId,
		Content:   data.Content,
		CreatedBy: userID,
		CreatedAt: time.Now(),
	}

	err = CreateComment(ctx, database.Client, threadID, c)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrCreateComment, CreateHandler)
		res.Message = WrapErrCreateComment.Message
		return res, err
	}

	res.Message = SuccessfulCreateCommentMessage

	return res, err
}

// AddKudo adds a new kudo for the given user and comment in the database
func AddKudo(ctx context.Context, client *ent.Client, userID uuid.UUID, commentID int) error {
	return client.CommentKudo.
		Create().
		SetUserID(userID).
		SetCommentID(commentID).
		Exec(ctx)
}

// HandleAddKudo parses the JSON body in the POST request, calls
// AddKudo if there is no error and returns a JSON encoded API response
func HandleAddKudo(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	userID, err := users.GetUserIDFromToken(r)
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrNotFound, CreateHandler)
		res.Message = customerrors.WrapErrNotFound.Message
		return res, err
	}

	commentID, err := strconv.Atoi(chi.URLParam(r, "commentID"))
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrStrToInt, AddKudoHandler)
		res.Message = customerrors.WrapErrStrToInt.Message
		return res, err
	}

	// Check if thread exists
	_, err = GetCommentById(ctx, commentID)
	if err != nil {
		switch t := err.(type) {
		default:
			res.Error = api.BuildError(t, WrapErrCheckCommentKudo, AddKudoHandler)
			res.Message = WrapErrCheckCommentKudo.Message
		case *ent.NotFoundError:
			res.Error = api.BuildError(t, customerrors.WrapErrNotFound, AddKudoHandler)
			res.Message = customerrors.WrapErrNotFound.Message
		}
		return res, err
	}

	// Check if user has already given a kudo to the same comment
	c, err := database.Client.CommentKudo.
		Query().
		Where(
			commentkudo.UserID(userID),
			commentkudo.CommentID(commentID),
		).
		Count(ctx)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrCheckCommentKudo, AddKudoHandler)
		res.Message = WrapErrCheckCommentKudo.Message
		return res, err
	}

	if c != 0 {
		res.Error = api.BuildError(ErrCommentKudoExist, WrapErrCommentKudoExist, AddKudoHandler)
		res.Message = ErrCommentKudoExist.Error()
		return res, ErrCommentKudoExist
	}

	err = AddKudo(ctx, database.Client, userID, commentID)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrAddKudo, AddKudoHandler)
		res.Message = WrapErrAddKudo.Message
		return res, err
	}

	res.Message = SuccessfulAddKudoMessage

	return res, err
}
