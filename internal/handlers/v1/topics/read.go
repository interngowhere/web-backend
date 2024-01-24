package topics

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/topic"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
)

const ReadHandler = "topics.HandleList"
const SuccessfulListTopicsMessage = "Listed all topics found"

// GetTopics returns the matching topic based on topic title slug provided
// If no slug is provided, all topics will be returned
func GetTopics(ctx context.Context, q string) ([]*ent.Topic, error) {
	if len(q) == 0 {
		return database.Client.Topic.
			Query().
			All(ctx)
	} else {
		return database.Client.Topic.
			Query().
			Where(topic.SlugEQ(q)).
			All(ctx)
	}
}

// ListByThread returns the corresponding topic of a given thread
func ListByThread(ctx context.Context, t *ent.Thread) (*ent.Topic, error) {
	return t.QueryTopics().Only(ctx)
}

// HandleList handles the GET request, calls
// GetTopics if needed and returns a JSON encoded API response
func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}
	data := []TopicResponse{}

	q := chi.URLParam(r, "title")
	topics, err := GetTopics(ctx, q)
	if err != nil {
		res = api.BuildError(err, WrapErrRetrieveTopics, ReadHandler)
		return res, err
	}

	for _, t := range topics {
		data = append(data, TopicResponse{
			ID:               t.ID,
			Title:            t.Title,
			Slug:             t.Slug,
			Description:      t.Description,
			ShortDescription: t.ShortDescription,
			ProfilePicURL:    t.ProfilePicURL,
			CreatedAt:        t.CreatedAt,
		})
	}

	if len(data) == 0 {
		res = api.BuildError(ErrNoTopicFound, customerrors.WrapErrNotFound, ReadHandler)
		return res, ErrNoTopicFound
	}

	encodedData, err := json.Marshal(data)
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrEncodeView, ReadHandler)
		res.Message = customerrors.WrapErrEncodeView.Message
	} else {
		res.Data = encodedData
		res.Message = SuccessfulListTopicsMessage
	}

	return res, err
}
