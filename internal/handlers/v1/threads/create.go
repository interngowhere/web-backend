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

	userID, err := users.GetUserIDFromToken(r)
	if err != nil {
		res = api.BuildError(err, users.WrapErrRetrieveIDFromJWT, CreateHandler)
		return res, err
	}

	topics, err := topics.GetTopics(ctx, chi.URLParam(r, "title"))
	if err != nil {
		res = api.BuildError(err, WrapErrRetrieveTopic, CreateHandler)
		return res, err
	}

	// Read JSON body in request into a new ThreadRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data ThreadRequest
	err = decoder.Decode(&data)
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
		return res, err
	}

	// Check for missing title in request body
	// Note: title is the only required field in the POST request
	if len(data.Title) == 0 {
		res = api.BuildError(customerrors.ErrMalformedRequest, customerrors.WrapErrRequestFormat, CreateHandler)
		return res, customerrors.ErrMalformedRequest
	}

	slug := mystr.ToLowerKebab(data.Title)

	t := ent.Thread{
		Title:       data.Title,
		Slug:        slug,
		Description: data.Description,
		CreatedBy:   userID,
	}

	err = CreateThread(ctx, database.Client, topics[0], t, data.Tags)
	if err != nil {
		res = api.BuildError(err, WrapErrCreateThread, CreateHandler)
		return res, err
	}

	res.Message = SuccessfulCreateThreadMessage
	res.Code = 201

	return res, err
}

// AddKudo adds a new kudo for the given user and thread in the database
func AddKudo(ctx context.Context, client *ent.Client, userID uuid.UUID, threadID int) error {
	return client.ThreadKudo.
		Create().
		SetUserID(userID).
		SetThreadID(threadID).
		Exec(ctx)
}

// HandleAddKudo parses the JSON body in the POST request, calls
// AddKudo if there is no error and returns a JSON encoded API response
func HandleAddKudo(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	userID, err := users.GetUserIDFromToken(r)
	if err != nil {
		res = api.BuildError(err, users.WrapErrRetrieveIDFromJWT, AddKudoHandler)
		return res, err
	}

	threadID, err := strconv.Atoi(chi.URLParam(r, "threadID"))
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrStrToInt, AddKudoHandler)
		return res, err
	}

	// Check if thread exists
	_, err = GetThreadByID(ctx, threadID)
	if err != nil {
		switch t := err.(type) {
		default:
			res = api.BuildError(t, WrapErrCheckThreadKudo, AddKudoHandler)
		case *ent.NotFoundError:
			res = api.BuildError(t, customerrors.WrapErrNotFound, AddKudoHandler)
		}
		return res, err
	}

	// Check if user has already given a kudo to the same thread
	c, err := database.Client.ThreadKudo.
		Query().
		Where(
			threadkudo.UserID(userID),
			threadkudo.ThreadID(threadID),
		).
		Count(ctx)
	if err != nil {
		res = api.BuildError(err, WrapErrCheckThreadKudo, AddKudoHandler)
		return res, err
	}

	if c != 0 {
		res = api.BuildError(ErrThreadKudoExist, WrapErrThreadKudoExist, AddKudoHandler)
		return res, ErrThreadKudoExist
	}

	err = AddKudo(ctx, database.Client, userID, threadID)
	if err != nil {
		res = api.BuildError(err, WrapErrAddKudo, AddKudoHandler)
		return res, err
	}

	res.Message = SuccessfulAddKudoMessage
	res.Code = 200

	return res, err
}
