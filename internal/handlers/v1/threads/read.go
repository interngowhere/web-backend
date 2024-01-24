package threads

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/thread"
	"github.com/interngowhere/web-backend/ent/threadkudo"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
	"github.com/interngowhere/web-backend/internal/handlers/v1/tags"
	"github.com/interngowhere/web-backend/internal/handlers/v1/topics"
	"github.com/interngowhere/web-backend/internal/handlers/v1/users"
)

const ReadHandler = "threads.HandleList"
const SuccessfulListThreadsMessage = "Listed all threads found"

// GetThreads returns the matching threads based on thread id
// provided. If no id is provided, all threads will be returned
func GetThreads(ctx context.Context, threadID int) ([]*ent.Thread, error) {
	if threadID == 0 {
		return database.Client.Thread.
			Query().
			All(ctx)
	} else {
		return database.Client.Thread.
			Query().
			Where(thread.ID(threadID)).
			All(ctx)
	}
}

func GetThreadByID(ctx context.Context, threadID int) (*ent.Thread, error) {
	return database.Client.Thread.
		Query().
		Where(thread.ID(threadID)).
		Only(ctx)
}

// GetTagsByThread returns all tags associated with a given thread
func GetTagsByThread(ctx context.Context, t *ent.Thread) ([]*ent.Tag, error) {
	return t.QueryTags().All(ctx)
}

// GetThreadKudoCount returns the number of kudos a thread has received
func GetThreadKudoCount(ctx context.Context, id int) (int, error) {
	return database.Client.ThreadKudo.
		Query().
		Where(threadkudo.ThreadID(id)).
		Count(ctx)
}

// CheckDidUserKudo checks if a user has given a kudo
// to a particular thread and returns a boolean
func CheckDidUserKudo(ctx context.Context, id int, userID uuid.UUID) (bool, error) {
	c, err := database.Client.ThreadKudo.
		Query().
		Where(threadkudo.ThreadID(id)).
		Where(threadkudo.UserID(userID)).
		Count(ctx)
	if err != nil {
		return false, err
	}
	return c == 1, err
}

// HandleList handles the GET request, calls
// GetThreads if needed and returns a JSON encoded API response
func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}
	data := []ThreadsResponse{}

	// get threads
	q := chi.URLParam(r, "threadID")
	threadID := 0

	if len(q) != 0 {
		threadID, _ = strconv.Atoi(q)
	}

	threads, err := GetThreads(ctx, threadID)
	if err != nil {
		res = api.BuildError(err, WrapErrRetrieveThreads, ReadHandler)
		return res, err
	}

	userID, _ := users.GetUserIDFromToken(r)

	for _, thread := range threads {
		t, err := GetTagsByThread(ctx, thread)
		if err != nil {
			res = api.BuildError(err, WrapErrRetrieveTags, ReadHandler)
			return res, err
		}

		formattedTags := []tags.TagResponse{}
		for _, tag := range t {
			formattedTags = append(formattedTags, tags.TagResponse{
				ID:      tag.ID,
				TagName: tag.TagName,
			})
		}

		c, err := GetThreadKudoCount(ctx, thread.ID)
		if err != nil {
			res = api.BuildError(err, WrapErrGetKudoCount, ReadHandler)
			return res, err
		}

		b, err := CheckDidUserKudo(ctx, thread.ID, userID)
		if err != nil {
			res = api.BuildError(err, WrapErrCheckDidUserKudo, ReadHandler)
			return res, err
		}

		u, err := users.GetUserFromID(ctx, database.Client, thread.CreatedBy)
		if err != nil {
			res = api.BuildError(err, users.WrapErrGetUserFromID, ReadHandler)
			return res, err
		}

		topic, err := topics.ListByThread(ctx, thread)
		if err != nil {
			res = api.BuildError(err, WrapErrGetTopicFromThread, ReadHandler)
			return res, err
		}

		data = append(data, ThreadsResponse{
			ID:                thread.ID,
			Title:             thread.Title,
			Slug:              thread.Slug,
			Description:       thread.Description,
			ModifiedAt:        thread.ModifiedAt,
			CreatedByID:       thread.CreatedBy,
			CreatedByUsername: u.Username,
			CreatedAt:         thread.CreatedAt,
			Tags:              &formattedTags,
			KudoCount:         c,
			UserKudoed:        b,
			TopicSlug:         topic.Slug,
		})
	}

	if len(data) == 0 {
		res = api.BuildError(ErrNoThreadFound, customerrors.WrapErrNotFound, ReadHandler)
		return res, ErrNoThreadFound
	}

	encodedData, err := json.Marshal(data)
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrEncodeView, ReadHandler)
	} else {
		res.Data = encodedData
		res.Message = SuccessfulListThreadsMessage
	}

	return res, err
}
