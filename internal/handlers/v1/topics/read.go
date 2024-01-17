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

// HandleList handles the GET request, calls
// GetTopics if needed and returns a JSON encoded API response
func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	res := &api.Response{}
	q := chi.URLParam(r, "title")
	topics, err := GetTopics(q)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveTopics, ReadHandler)
		res.Message = WrapErrRetrieveTopics.Message
		return res, err
	}

	data, err := json.Marshal(topics)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrEncodeView, ReadHandler)
		res.Message = WrapErrEncodeView.Message
	} else {
		res.Data = data
		res.Message = SuccessfulListTopicsMessage
	}

	return res, err
}
