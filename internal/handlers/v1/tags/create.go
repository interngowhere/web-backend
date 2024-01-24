package tags

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/tag"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
)

const CreateHandler = "tags.HandleCreate"
const SuccessfulCreateTagMessage = "Successfully created tag"

// CreateTag adds a new tag entry in the database
func CreateTag(ctx context.Context, client *ent.Client, tag ent.Tag) error {
	return client.Tag.
		Create().
		SetTagName(tag.TagName).
		Exec(ctx)
}

// HandleCreate parses the POST request form data, calls
// CreateTag if needed and returns a JSON encoded API response
func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	// Read JSON body in request into a new TagRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data TagRequest
	err := decoder.Decode(&data)
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
		return res, err
	}

	// Check for missing input fields in request body
	if len(data.TagName) == 0 {
		res = api.BuildError(customerrors.ErrMalformedRequest, customerrors.WrapErrRequestFormat, CreateHandler)
		return res, customerrors.ErrMalformedRequest
	}

	// Check if tag already exist
	c, err := database.Client.Tag.
		Query().
		Where(tag.TagName(data.TagName)).
		Count(ctx)
	if err != nil {
		res = api.BuildError(err, WrapErrCheckTagExist, CreateHandler)
		return res, err
	}
	if c != 0 {
		res = api.BuildError(ErrTagExist, customerrors.WrapErrResourceExist, CreateHandler)
		return res, ErrTagExist
	}

	t := ent.Tag{
		TagName: data.TagName,
	}

	err = CreateTag(ctx, database.Client, t)
	if err != nil {
		res = api.BuildError(err, WrapErrCreateTag, ReadHandler)
		return res, err
	}

	res.Message = SuccessfulCreateTagMessage
	res.Code = 201

	return res, err
}
