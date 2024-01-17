package tags

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
)

const DeleteHandler = "tags.HandleDelete"
const SuccessfulDeleteTagMessage = "Successfully deleted tag"

// DeleteTag deletes an existing tag entry in the database
// based on ent.Tag object provided
func DeleteTag(ctx context.Context, client *ent.Client, tagId int) error {
	return client.Tag.
		DeleteOneID(tagId).
		Exec(ctx)
}

// HandleDelete handles the DELETE request, calls
// DeleteTag if needed and returns a JSON encoded API response
func HandleDelete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	tagId, err := strconv.Atoi(chi.URLParam(r, "tagId"))
	if err != nil {
		res.Error = api.BuildError(err, WrapErrStrToInt, DeleteHandler)
		res.Message = WrapErrStrToInt.Message
		return res, err
	}

	t, err := GetTagByID(tagId)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveTags, DeleteHandler)
		res.Message = WrapErrRetrieveTags.Message
		return res, err
	}

	err = DeleteTag(ctx, database.Client, t.ID)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrDeleteTag, DeleteHandler)
		res.Message = WrapErrDeleteTag.Message
		return res, err
	}

	res.Message = SuccessfulDeleteTagMessage

	return res, err
}
