package topics

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
)

const UpdateHandler = "topics.HandleUpdate"
const SuccessfulUpdateTopicMessage = "Successfully updated topic"

// UpdateTopic updates an existing topic entry in the database
// based on topic ID provided via the ent.Topic object
func UpdateTopic(ctx context.Context, client *ent.Client, topic *ent.Topic) error {
	return client.Topic.
		UpdateOne(topic).
		SetDescription(topic.Description).
		SetShortDescription(topic.ShortDescription).
		SetProfilePicURL(topic.ProfilePicURL).
		Exec(ctx)
}

// HandleCreate parses the PUT request form data, calls
// UpdateTopic if needed and returns a JSON encoded API response
func HandleUpdate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	slug := chi.URLParam(r, "title")

	// Retrieve a reference to the topic
	t, err := GetTopics(ctx, slug)
	if err != nil {
		res = api.BuildError(err, WrapErrRetrieveTopics, UpdateHandler)
		return res, err
	}
	if len(t) == 0 {
		res = api.BuildError(customerrors.ErrResourceNotFound, customerrors.WrapErrNotFound, UpdateHandler)
		res.Message = customerrors.WrapErrNotFound.Message
		return res, customerrors.ErrResourceNotFound
	}

	// Read JSON body in request into a new TopicRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data TopicRequest
	err = decoder.Decode(&data)
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
		return res, err
	}

	// Update fields in topic object only if a new value for the field is provided
	if len(data.Description) != 0 {
		t[0].Description = data.Description
	}
	if len(data.ShortDescription) != 0 {
		t[0].ShortDescription = data.ShortDescription
	}
	if len(data.ProfilePicURL) != 0 {
		t[0].ProfilePicURL = data.ProfilePicURL
	}

	err = UpdateTopic(ctx, database.Client, t[0])
	if err != nil {
		res = api.BuildError(err, WrapErrUpdateTopic, UpdateHandler)
		return res, err
	}

	res.Message = SuccessfulUpdateTopicMessage

	return res, err
}
