package topics

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/topic"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
	mystr "github.com/interngowhere/web-backend/internal/utils"
)

const CreateHandler = "topics.HandleCreate"
const SuccessfulCreateTopicMessage = "Successfully created topic"

// CreateTopic adds a new topic entry in the database
func CreateTopic(ctx context.Context, client *ent.Client, topic ent.Topic) error {
	return client.Topic.
		Create().
		SetTitle(topic.Title).
		SetSlug(topic.Slug).
		SetDescription(topic.Description).
		SetShortDescription(topic.ShortDescription).
		SetProfilePicURL(topic.ProfilePicURL).
		SetCreatedAt(time.Now()).
		Exec(ctx)
}

// HandleCreate parses the POST request form data, calls
// CreateTopic if needed and returns a JSON encoded API response
func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	// Read JSON body in request into a new TopicRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data TopicRequest
	err := decoder.Decode(&data)
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
		return res, err
	}

	// Check for missing input fields in request body
	if len(data.Title) == 0 ||
		len(data.Description) == 0 ||
		len(data.ShortDescription) == 0 {
		res = api.BuildError(customerrors.ErrMalformedRequest, customerrors.WrapErrRequestFormat, CreateHandler)
		return res, customerrors.ErrMalformedRequest
	}

	slug := mystr.ToLowerKebab(data.Title)

	// Check if topic already exist
	c, err := database.Client.Topic.
		Query().
		Where(
			topic.
				Title(data.Title),
		).
		Count(ctx)
	if err != nil {
		res = api.BuildError(err, WrapErrCheckTopicExist, CreateHandler)
		return res, err
	}
	if c != 0 {
		res = api.BuildError(ErrTopicExist, customerrors.WrapErrResourceExist, CreateHandler)
		return res, ErrTopicExist
	}

	t := ent.Topic{
		Title:            data.Title,
		Slug:             slug,
		Description:      data.Description,
		ShortDescription: data.ShortDescription,
		ProfilePicURL:    data.ProfilePicURL,
	}

	err = CreateTopic(ctx, database.Client, t)
	if err != nil {
		res = api.BuildError(err, WrapErrCreateTopic, ReadHandler)
		return res, err
	}

	res.Message = SuccessfulCreateTopicMessage
	res.Code = 201

	return res, err
}
