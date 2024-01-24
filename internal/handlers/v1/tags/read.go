package tags

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/tag"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
)

const ReadHandler = "tags.HandleList"
const SuccessfulListTagsMessage = "Listed all tags found"

// GetTags returns a list of matching tags based on the query parameter provided.
// If no query parameter is provided, all tags will be returned by default.
func GetTags(q string) ([]*ent.Tag, error) {
	return database.Client.Tag.
		Query().
		Where(tag.TagNameContains(q)).
		All(context.Background())
}

// GetTagByID returns a single tag based on the tag ID provided
func GetTagByID(tagId int) (*ent.Tag, error) {
	return database.Client.Tag.
		Query().
		Where(tag.ID(tagId)).
		Only(context.Background())
}

// HandleList handles the GET request, calls
// GetTopics if needed and returns a JSON encoded API response
func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	res := &api.Response{}
	q := r.URL.Query().Get("name")
	q = strings.ToLower(q)
	t, err := GetTags(q)
	if err != nil {
		res = api.BuildError(err, WrapErrRetrieveTags, ReadHandler)
		return res, err
	}

	data := []TagResponse{}
	for _, tag := range t {
		data = append(data, TagResponse{
			ID:      tag.ID,
			TagName: tag.TagName,
		})
	}

	if len(data) == 0 {
		res = api.BuildError(ErrNoTagFound, customerrors.WrapErrNotFound, ReadHandler)
		return res, ErrNoTagFound
	}

	encodedData, err := json.Marshal(data)
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrEncodeView, ReadHandler)
		res.Message = customerrors.WrapErrEncodeView.Message
	} else {
		res.Data = encodedData
		res.Message = SuccessfulListTagsMessage
	}

	return res, err
}
