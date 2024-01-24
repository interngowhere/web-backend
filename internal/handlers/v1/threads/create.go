package threads

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/threadkudo"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
	"github.com/interngowhere/web-backend/internal/handlers/v1/topics"
	"github.com/interngowhere/web-backend/internal/handlers/v1/users"
	mystr "github.com/interngowhere/web-backend/internal/utils"
)

const CreateHandler = "threads.HandleCreate"
const AddKudoHandler = "threads.HandleAddKudo"
const SuccessfulCreateThreadMessage = "Successfully created thread"
const SuccessfulAddKudoMessage = "Successfully added kudo to thread"

// CreateThread adds a new thread entry in the database
func CreateThread(ctx context.Context, client *ent.Client, topic *ent.Topic, thread ent.Thread, tagIDs []int) error {
	return client.Thread.
		Create().
		SetTitle(thread.Title).
		SetSlug(thread.Slug).
		SetDescription(thread.Description).
		SetCreatedBy(thread.CreatedBy).
		SetTopics(topic).
		SetCreatedAt(time.Now()).
		AddTagIDs(tagIDs...).
		Exec(ctx)
}

// HandleCreate parses the POST request form data, calls
// CreateThread if needed and returns a JSON encoded API response
func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	userId, err := users.GetUserIDFromToken(r)
	if err != nil {
		res.Error = api.BuildError(err, users.WrapErrRetrieveIDFromJWT, CreateHandler)
		res.Message = users.WrapErrRetrieveIDFromJWT.Message
		return res, err
	}

	topics, err := topics.GetTopics(ctx, chi.URLParam(r, "title"))
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveTopic, CreateHandler)
		res.Message = WrapErrRetrieveTopic.Message
		return res, err
	}

	// Read JSON body in request into a new ThreadRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data ThreadRequest
	err = decoder.Decode(&data)
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
		res.Message = customerrors.WrapErrDecodeRequest.Message
		return res, err
	}

	// Check for missing title in request body
	// Note: title is the only required field in the POST request
	if len(data.Title) == 0 {
		res.Error = api.BuildError(customerrors.ErrMalformedRequest, customerrors.WrapErrRequestFormat, CreateHandler)
		res.Message = customerrors.WrapErrRequestFormat.Message
		return res, customerrors.ErrMalformedRequest
	}

	slug := mystr.ToLowerKebab(data.Title)

	t := ent.Thread{
		Title:       data.Title,
		Slug:        slug,
		Description: data.Description,
		CreatedBy:   userId,
	}

	err = CreateThread(ctx, database.Client, topics[0], t, data.Tags)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrCreateThread, CreateHandler)
		res.Message = WrapErrCreateThread.Message
		return res, err
	}

	res.Message = SuccessfulCreateThreadMessage

	return res, err
}

// AddKudo adds a new kudo for the given user and thread in the database
func AddKudo(ctx context.Context, client *ent.Client, userId uuid.UUID, threadId int) error {
	return client.ThreadKudo.
		Create().
		SetUserID(userId).
		SetThreadID(threadId).
		Exec(ctx)
}

// HandleAddKudo parses the JSON body in the POST request, calls
// AddKudo if there is no error and returns a JSON encoded API response
func HandleAddKudo(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	userId, err := users.GetUserIDFromToken(r)
	if err != nil {
		res.Error = api.BuildError(err, users.WrapErrRetrieveIDFromJWT, AddKudoHandler)
		res.Message = users.WrapErrRetrieveIDFromJWT.Message
		return res, err
	}

	threadId, err := strconv.Atoi(chi.URLParam(r, "threadId"))
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrStrToInt, AddKudoHandler)
		res.Message = customerrors.WrapErrStrToInt.Message
		return res, err
	}

	// Check if thread exists
	_, err = GetThreadByID(ctx, threadId)
	if err != nil {
		switch t := err.(type) {
		default:
			res.Error = api.BuildError(t, WrapErrCheckThreadKudo, AddKudoHandler)
			res.Message = WrapErrCheckThreadKudo.Message
		case *ent.NotFoundError:
			res.Error = api.BuildError(t, customerrors.WrapErrNotFound, AddKudoHandler)
			res.Message = customerrors.WrapErrNotFound.Message
		}
		return res, err
	}

	// Check if user has already given a kudo to the same thread
	c, err := database.Client.ThreadKudo.
		Query().
		Where(
			threadkudo.UserID(userId),
			threadkudo.ThreadID(threadId),
		).
		Count(ctx)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrCheckThreadKudo, AddKudoHandler)
		res.Message = WrapErrCheckThreadKudo.Message
		return res, err
	}

	if c != 0 {
		res.Error = api.BuildError(ErrThreadKudoExist, WrapErrThreadKudoExist, AddKudoHandler)
		res.Message = ErrThreadKudoExist.Error()
		return res, ErrThreadKudoExist
	}

	err = AddKudo(ctx, database.Client, userId, threadId)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrAddKudo, AddKudoHandler)
		res.Message = WrapErrAddKudo.Message
		return res, err
	}

	res.Message = SuccessfulCreateThreadMessage

	return res, err
}
