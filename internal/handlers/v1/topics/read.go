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
)

const ReadHandler = "topics.HandleList"
const SuccessfulListTopicsMessage = "Listed all topics found"

// GetTopics returns the matching topic based on topic title slug provided
// If no slug is provided, all topics will be returned
func GetTopics(q string) ([]*ent.Topic, error) {
	if len(q) == 0 {
		return database.Client.Topic.
			Query().
			All(context.Background())
	} else {
		return database.Client.Topic.
			Query().
			Where(topic.SlugEQ(q)).
			All(context.Background())
	}
}

// ListByThread returns the corresponding topic of a given thread
func ListByThread(t *ent.Thread) (*ent.Topic, error) {
	return t.QueryTopics().Only(context.Background())
}

// HandleList handles the GET request, calls
// GetTopics if needed and returns a JSON encoded API response
func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	res := &api.Response{}
	data := []TopicResponse{}

	q := chi.URLParam(r, "title")
	topics, err := GetTopics(q)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveTopics, ReadHandler)
		res.Message = WrapErrRetrieveTopics.Message
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

	encodedData, err := json.Marshal(data)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrEncodeView, ReadHandler)
		res.Message = WrapErrEncodeView.Message
	} else {
		res.Data = encodedData
		res.Message = SuccessfulListTopicsMessage
	}

	return res, err
}
