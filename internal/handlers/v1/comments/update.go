package comments

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
)

const UpdateHandler = "comments.HandleUpdate"
const SuccessfulUpdateCommentMessage = "Successfully updated comment"

// UpdateComment updates an existing comment entry in the database
// based on comment ID provided via the ent.Comment object
func UpdateComment(ctx context.Context, client *ent.Client, comment *ent.Comment) error {
	return client.Comment.
		UpdateOneID(comment.ID).
		SetContent(comment.Content).
		SetModifiedAt(time.Now()).
		Exec(ctx)
}

// HandleUpdate parses the PUT request form data, calls
// UpdateThread if needed and returns a JSON encoded API response
func HandleUpdate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	// Retrieve a reference to the thread
	commentID, err := strconv.Atoi(chi.URLParam(r, "commentID"))
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrStrToInt, ListHandler)
		return res, err
	}

	c, err := GetCommentById(ctx, commentID)
	if err != nil {
		res = api.BuildError(err, WrapErrRetrieveComments, UpdateHandler)
		return res, err
	}
	if c == nil {
		res = api.BuildError(ErrNoMatchFromCommentID, customerrors.WrapErrNotFound, UpdateHandler)
		return res, ErrNoMatchFromCommentID
	}

	// Read JSON body in request into a new ThreadRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data CommentRequest
	err = decoder.Decode(&data)
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrDecodeRequest, UpdateHandler)
		return res, err
	}

	// Update fields in thread object only if a new value for the field is provided
	if len(data.Content) != 0 {
		c.Content = data.Content
	}

	err = UpdateComment(ctx, database.Client, c)
	if err != nil {
		res = api.BuildError(err, WrapErrUpdateComment, UpdateHandler)
		return res, err
	}

	res.Message = SuccessfulUpdateCommentMessage

	return res, err
}
