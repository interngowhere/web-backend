package comments

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/comment"
	"github.com/interngowhere/web-backend/ent/commentkudo"
	"github.com/interngowhere/web-backend/ent/thread"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
	"github.com/interngowhere/web-backend/internal/handlers/v1/users"
)

const ListHandler = "comments.HandleList"
const SuccessfulListCommentsMessage = "Listed all comments found"

// GetRootComments return a list of root comments (comments
// which do not have a parent) in a given thread
func GetRootComments(ctx context.Context, threadId int) ([]*ent.Comment, error) {
	return database.Client.Comment.
		Query().
		Where(
			comment.HasThreadsWith(
				thread.ID(threadId),
			),
			comment.ParentID(0),
		).
		All(ctx)
}

// GetCommentById returns a single comment matching the ID provided
func GetCommentById(ctx context.Context, commentId int) (*ent.Comment, error) {
	return database.Client.Comment.
		Query().
		Where(comment.ID(commentId)).
		Only(ctx)
}

// GetCommentsByParent returns a list of children comments
// which are nested under the parent comment
func GetCommentsByParent(ctx context.Context, parentId int) ([]*ent.Comment, error) {
	return database.Client.Comment.
		Query().
		Where(comment.ParentID(parentId)).
		All(ctx)
}

func GetCommentKudoCount(ctx context.Context, id int) (int, error) {
	return database.Client.CommentKudo.
		Query().
		Where(commentkudo.CommentID(id)).
		Count(ctx)
}

func CheckDidUserKudo(ctx context.Context, id int, userId uuid.UUID) (bool, error) {
	c, err := database.Client.CommentKudo.
		Query().
		Where(commentkudo.CommentID(id)).
		Where(commentkudo.UserID(userId)).
		Count(ctx)
	if err != nil {
		return false, err
	}
	return c == 1, err
}

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}
	data := []CommentsResponse{}
	threadId, err := strconv.Atoi(chi.URLParam(r, "threadId"))
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrStrToInt, ListHandler)
		res.Message = customerrors.WrapErrStrToInt.Message
		return res, err
	}

	parents, err := GetRootComments(ctx, threadId)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveComments, ListHandler)
		res.Message = WrapErrRetrieveComments.Message
		return res, err
	}

	userId, _ := users.GetUserIDFromToken(r)

	for _, parent := range parents {
		children, err := GetCommentsByParent(ctx, parent.ID)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrRetrieveComments, ListHandler)
			res.Message = WrapErrRetrieveComments.Message
			return res, err
		}

		formattedChildren := []CommentsResponse{}
		for _, child := range children {
			c, err := GetCommentKudoCount(ctx, child.ID)
			if err != nil {
				res.Error = api.BuildError(err, WrapErrGetKudoCount, ListHandler)
				res.Message = WrapErrGetKudoCount.Message
				return res, err
			}

			b, err := CheckDidUserKudo(ctx, child.ID, userId)
			if err != nil {
				res.Error = api.BuildError(err, WrapErrCheckDidUserKudo, ListHandler)
				res.Message = WrapErrCheckDidUserKudo.Message
				return res, err
			}

			u, err := users.GetUsernameFromID(ctx, database.Client, child.CreatedBy)
			if err != nil {
				res.Error = api.BuildError(err, WrapErrGetUsernameFromID, ListHandler)
				res.Message = WrapErrGetUsernameFromID.Message
				return res, err
			}

			formattedChildren = append(formattedChildren, CommentsResponse{
				ID:                child.ID,
				ParentID:          child.ParentID,
				Content:           child.Content,
				ModifiedAt:        child.ModifiedAt,
				CreatedByID:       child.CreatedBy,
				CreatedByUsername: u,
				CreatedAt:         child.CreatedAt,
				KudoCount:         c,
				UserKudoed:        b,
			})
		}

		c, err := GetCommentKudoCount(ctx, parent.ID)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrGetKudoCount, ListHandler)
			res.Message = WrapErrGetKudoCount.Message
			return res, err
		}

		b, err := CheckDidUserKudo(ctx, parent.ID, userId)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrCheckDidUserKudo, ListHandler)
			res.Message = WrapErrCheckDidUserKudo.Message
			return res, err
		}

		u, err := users.GetUsernameFromID(ctx, database.Client, parent.CreatedBy)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrGetUsernameFromID, ListHandler)
			res.Message = WrapErrGetUsernameFromID.Message
			return res, err
		}

		data = append(data, CommentsResponse{
			ID:                parent.ID,
			ParentID:          parent.ParentID,
			Content:           parent.Content,
			ModifiedAt:        parent.ModifiedAt,
			CreatedByID:       parent.CreatedBy,
			CreatedByUsername: u,
			CreatedAt:         parent.CreatedAt,
			KudoCount:         c,
			UserKudoed:        b,
			Children:          &formattedChildren,
		})
	}

	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveComments, ListHandler)
		res.Message = WrapErrRetrieveComments.Message
		return res, err
	}

	encodedData, err := json.Marshal(data)
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrEncodeView, ListHandler)
		res.Message = customerrors.WrapErrEncodeView.Message
	} else {
		res.Data = encodedData
		res.Message = SuccessfulListCommentsMessage
	}

	return res, err
}
