package topics

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
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
	t, err := GetTopics(ctx, slug)
	if err != nil {
		res = api.BuildError(err, WrapErrRetrieveTopics, DeleteHandler)
		return res, err
	}
	if len(t) == 0 {
		res = api.BuildError(customerrors.ErrResourceNotFound, customerrors.WrapErrNotFound, DeleteHandler)
		return res, customerrors.ErrResourceNotFound
	}

	err = DeleteTopic(ctx, database.Client, t[0])
	if err != nil {
		res = api.BuildError(err, WrapErrDeleteTopic, DeleteHandler)
		return res, err
	}

	res.Message = SuccessfulDeleteTopicMessage

	return res, err
}
