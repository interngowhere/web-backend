package topics

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
)

const DeleteHandler = "topics.HandleDelete"
const SuccessfulDeleteTopicMessage = "Successfully deleted topic"

// DeleteTopic deletes an existing topic entry in the database
// based on ent.Topic object provided
func DeleteTopic(ctx context.Context, client *ent.Client, topic *ent.Topic) error {
	return client.Topic.
		DeleteOne(topic).
		Exec(ctx)
}

// HandleDelete handles the DELETE request, calls
// DeleteTopic if needed and returns a JSON encoded API response
func HandleDelete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	slug := chi.URLParam(r, "title")
	t, err := GetTopics(slug)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveTopics, DeleteHandler)
		res.Message = WrapErrRetrieveTopics.Message
		return res, err
	}
	if len(t) == 0 {
		res.Error = api.BuildError(ErrNoMatchFromSlug, WrapErrNoTopicFound, DeleteHandler)
		res.Message = WrapErrNoTopicFound.Message
		return res, ErrNoMatchFromSlug
	}

	err = DeleteTopic(ctx, database.Client, t[0])
	if err != nil {
		res.Error = api.BuildError(err, WrapErrDeleteTopic, DeleteHandler)
		res.Message = WrapErrDeleteTopic.Message
		return res, err
	}

	res.Message = SuccessfulDeleteTopicMessage

	return res, err
}
