package threads

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
	mystr "github.com/interngowhere/web-backend/internal/utils"
)

const UpdateHandler = "threads.HandleUpdate"
const SuccessfulUpdateThreadMessage = "Successfully updated thread"

// UpdateThread updates an existing thread entry in the database
// based on thread ID provided via the ent.Thread object
func UpdateThread(ctx context.Context, client *ent.Client, thread *ent.Thread, tagIDs []int) error {
	return client.Thread.
		UpdateOne(thread).
		SetTitle(thread.Title).
		SetSlug(thread.Slug).
		SetDescription(thread.Description).
		SetModifiedAt(time.Now()).
		ClearTags().
		AddTagIDs(tagIDs...).
		Exec(ctx)
}

// HandleUpdate parses the PUT request form data, calls
// UpdateThread if needed and returns a JSON encoded API response
func HandleUpdate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	// Retrieve a reference to the thread
	threadID, err := strconv.Atoi(chi.URLParam(r, "threadID"))
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrStrToInt, UpdateHandler)
		res.Message = customerrors.WrapErrStrToInt.Message
		return res, err
	}
	t, err := GetThreadByID(ctx, threadID)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveThreads, UpdateHandler)
		res.Message = WrapErrRetrieveThreads.Message
		return res, err
	}
	if t == nil {
		res.Error = api.BuildError(customerrors.ErrResourceNotFound, customerrors.WrapErrNotFound, UpdateHandler)
		res.Message = customerrors.WrapErrNotFound.Message
		return res, customerrors.ErrResourceNotFound
	}

	// Read JSON body in request into a new ThreadRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data ThreadRequest
	err = decoder.Decode(&data)
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrDecodeRequest, UpdateHandler)
		res.Message = customerrors.WrapErrDecodeRequest.Message
		return res, err
	}

	// Update fields in thread object only if a new value for the field is provided
	if len(data.Title) != 0 {
		t.Title = data.Title
		t.Slug = mystr.ToLowerKebab(data.Title)
	}
	if len(data.Description) != 0 {
		t.Description = data.Description
	}

	err = UpdateThread(ctx, database.Client, t, data.Tags)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrUpdateThread, UpdateHandler)
		res.Message = WrapErrUpdateThread.Message
		return res, err
	}

	res.Message = SuccessfulUpdateThreadMessage

	return res, err
}
