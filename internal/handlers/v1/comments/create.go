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
	"github.com/interngowhere/web-backend/internal/handlers/v1/threads"
	"github.com/interngowhere/web-backend/internal/handlers/v1/users"
)

const CreateHandler = "comments.HandleCreate"
const AddKudoHandler = "comments.HandleAddKudo"
const SuccessfulCreateCommentMessage = "Successfully created comment"
const SuccessfulAddKudoMessage = "Successfully added kudo to comment"

// CreateComment adds a new comment entry in the database
func CreateComment(ctx context.Context, client *ent.Client, threadId int, comment ent.Comment) error {
	return client.Comment.
		Create().
		SetParentID(comment.ParentID).
		SetContent(comment.Content).
		SetCreatedBy(comment.CreatedBy).
		SetCreatedAt(comment.CreatedAt).
		SetCommentAuthorsID(comment.CreatedBy).
		SetThreadsID(threadId).
		Exec(ctx)
}

// HandleCreate parses the POST request form data, calls
// CreateComment if needed and returns a JSON encoded API response
func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	userId, err := users.GetUserIDFromToken(r)
	if err != nil {
		res.Error = api.BuildError(err, users.WrapErrRetrieveUserID, CreateHandler)
		res.Message = users.WrapErrRetrieveUserID.Message
		return res, err
	}

	threadId, err := strconv.Atoi(chi.URLParam(r, "threadId"))
	if err != nil {
		res.Error = api.BuildError(err, WrapErrStrToInt, ListHandler)
		res.Message = WrapErrStrToInt.Message
		return res, err
	}

	_, err = threads.GetThreadByID(threadId)
	if err != nil {
		res.Error = api.BuildError(err, threads.WrapErrRetrieveThreads, CreateHandler)
		res.Message = threads.WrapErrRetrieveThreads.Message
		return res, err
	}
	// if len(t) == 0 {
	// 	res.Error = api.BuildError(ErrNoMatchFromThreadID, threads.WrapErrRetrieveThreads, CreateHandler)
	// 	res.Message = threads.WrapErrRetrieveThreads.Message
	// 	return res, ErrNoMatchFromThreadID
	// }

	// Read JSON body in request into a new ThreadRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data CommentRequest
	err = decoder.Decode(&data)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrDecodeRequest, CreateHandler)
		res.Message = WrapErrDecodeRequest.Message
		return res, err
	}

	// Check if content field is present in request body
	// Note: content is the only required field in the POST request
	// If parentId field is not provided, parentId will be set to 0
	// by default, i.e., the comment has no parent.
	if len(data.Content) == 0 {
		res.Error = api.BuildError(ErrMissingContent, WrapErrRequestFormat, CreateHandler)
		res.Message = WrapErrRequestFormat.Message
		return res, ErrMissingContent
	}

	if len(data.ParentID) == 0 {
		data.ParentID = "0"
	}
	parentId, err := strconv.Atoi(data.ParentID)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrStrToInt, CreateHandler)
		res.Message = WrapErrStrToInt.Message
		return res, err
	}

	c := ent.Comment{
		ParentID:  parentId,
		Content:   data.Content,
		CreatedBy: userId,
		CreatedAt: time.Now(),
	}

	err = CreateComment(ctx, database.Client, threadId, c)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrCreateComment, CreateHandler)
		res.Message = WrapErrCreateComment.Message
		return res, err
	}

	res.Message = SuccessfulCreateCommentMessage

	return res, err
}

// AddKudo adds a new kudo for the given user and comment in the database
func AddKudo(ctx context.Context, client *ent.Client, userId uuid.UUID, commentId int) error {
	return client.CommentKudo.
		Create().
		SetUserID(userId).
		SetCommentID(commentId).
		Exec(ctx)
}

// HandleAddKudo parses the JSON body in the POST request, calls
// AddKudo if there is no error and returns a JSON encoded API response
func HandleAddKudo(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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

	// Check if thread exists
	_, err = GetCommentById(commentId)
	if err != nil {
		switch t := err.(type) {
		default:
			res.Error = api.BuildError(t, WrapErrCheckCommentKudo, AddKudoHandler)
			res.Message = WrapErrCheckCommentKudo.Message
		case *ent.NotFoundError:
			res.Error = api.BuildError(t, WrapErrNoCommentFound, AddKudoHandler)
			res.Message = WrapErrNoCommentFound.Message
		}
		return res, err
	}

	// Check if user has already given a kudo to the same comment
	c, err := database.Client.CommentKudo.
		Query().
		Where(
			commentkudo.UserID(userId),
			commentkudo.CommentID(commentId),
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

	err = AddKudo(ctx, database.Client, userId, commentId)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrAddKudo, AddKudoHandler)
		res.Message = WrapErrAddKudo.Message
		return res, err
	}

	res.Message = SuccessfulAddKudoMessage

	return res, err
}
