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
	"github.com/interngowhere/web-backend/internal/handlers/v1/users"
)

const ListHandler = "comments.HandleList"
const SuccessfulListCommentsMessage = "Listed all comments found"

// GetRootComments return a list of root comments (comments
// which do not have a parent) in a given thread
func GetRootComments(threadId int) ([]*ent.Comment, error) {
	return database.Client.Comment.
		Query().
		Where(
			comment.HasThreadsWith(
				thread.ID(threadId),
			),
			comment.ParentID(0),
		).
		All(context.Background())
}

// GetCommentById returns a single comment matching the ID provided
func GetCommentById(commentId int) (*ent.Comment, error) {
	return database.Client.Comment.
		Query().
		Where(comment.ID(commentId)).
		Only(context.Background())
}

// GetCommentsByParent returns a list of children comments
// which are nested under the parent comment
func GetCommentsByParent(parentId int) ([]*ent.Comment, error) {
	return database.Client.Comment.
		Query().
		Where(comment.ParentID(parentId)).
		All(context.Background())
}

func GetCommentKudoCount(id int) (int, error) {
	return database.Client.CommentKudo.
		Query().
		Where(commentkudo.CommentID(id)).
		Count(context.Background())
}

func CheckDidUserKudo(id int, userId uuid.UUID) (bool, error) {
	c, err := database.Client.CommentKudo.
		Query().
		Where(commentkudo.CommentID(id)).
		Where(commentkudo.UserID(userId)).
		Count(context.Background())
	if err != nil {
		return false, err
	}
	return c == 1, err
}

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	res := &api.Response{}
	data := []CommentsResponse{}
	threadId, err := strconv.Atoi(chi.URLParam(r, "threadId"))
	if err != nil {
		res.Error = api.BuildError(err, WrapErrStrToInt, ListHandler)
		res.Message = WrapErrStrToInt.Message
		return res, err
	}

	parents, err := GetRootComments(threadId)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveComments, ListHandler)
		res.Message = WrapErrRetrieveComments.Message
		return res, err
	}

	userId, err := users.GetUserIDFromToken(r)
	if err != nil {
		res.Error = api.BuildError(err, users.WrapErrRetrieveUserID, ListHandler)
		res.Message = users.WrapErrRetrieveUserID.Message
		return res, err
	}

	for _, parent := range parents {
		children, err := GetCommentsByParent(parent.ID)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrRetrieveComments, ListHandler)
			res.Message = WrapErrRetrieveComments.Message
			return res, err
		}

		formattedChildren := []CommentsResponse{}
		for _, child := range children {
			c, err := GetCommentKudoCount(child.ID)
			if err != nil {
				res.Error = api.BuildError(err, WrapErrGetKudoCount, ListHandler)
				res.Message = WrapErrGetKudoCount.Message
				return res, err
			}

			b, err := CheckDidUserKudo(child.ID, userId)
			if err != nil {
				res.Error = api.BuildError(err, WrapErrCheckDidUserKudo, ListHandler)
				res.Message = WrapErrCheckDidUserKudo.Message
				return res, err
			}

			formattedChildren = append(formattedChildren, CommentsResponse{
				ID:         child.ID,
				ParentID:   child.ParentID,
				Content:    child.Content,
				ModifiedAt: child.ModifiedAt,
				CreatedBy:  child.CreatedBy,
				CreatedAt:  child.CreatedAt,
				KudoCount:  c,
				UserKudoed: b,
			})
		}

		c, err := GetCommentKudoCount(parent.ID)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrGetKudoCount, ListHandler)
			res.Message = WrapErrGetKudoCount.Message
			return res, err
		}

		b, err := CheckDidUserKudo(parent.ID, userId)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrCheckDidUserKudo, ListHandler)
			res.Message = WrapErrCheckDidUserKudo.Message
			return res, err
		}

		data = append(data, CommentsResponse{
			ID:         parent.ID,
			ParentID:   parent.ParentID,
			Content:    parent.Content,
			ModifiedAt: parent.ModifiedAt,
			CreatedBy:  parent.CreatedBy,
			CreatedAt:  parent.CreatedAt,
			KudoCount:  c,
			UserKudoed: b,
			Children:   &formattedChildren,
		})
	}

	if err != nil {
		res.Error = api.BuildError(err, WrapErrRetrieveComments, ListHandler)
		res.Message = WrapErrRetrieveComments.Message
		return res, err
	}

	encodedData, err := json.Marshal(data)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrEncodeView, ListHandler)
		res.Message = WrapErrEncodeView.Message
	} else {
		res.Data = encodedData
		res.Message = SuccessfulListCommentsMessage
	}

	return res, err
}
