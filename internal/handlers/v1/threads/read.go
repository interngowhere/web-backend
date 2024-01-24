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
	"github.com/interngowhere/web-backend/internal/handlers/v1/tags"
	"github.com/interngowhere/web-backend/internal/handlers/v1/topics"
	"github.com/interngowhere/web-backend/internal/handlers/v1/users"
)

const ReadHandler = "threads.HandleList"
const SuccessfulListThreadsMessage = "Listed all threads found"

// GetThreads returns the matching threads based on thread id
// provided. If no id is provided, all threads will be returned
func GetThreads(ctx context.Context, threadId int) ([]*ent.Thread, error) {
	if threadId == 0 {
		return database.Client.Thread.
			Query().
			All(ctx)
	} else {
		return database.Client.Thread.
			Query().
			Where(thread.ID(threadId)).
			All(ctx)
	}
}

func GetThreadByID(ctx context.Context, threadId int) (*ent.Thread, error) {
	return database.Client.Thread.
		Query().
		Where(thread.ID(threadId)).
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
func CheckDidUserKudo(ctx context.Context, id int, userId uuid.UUID) (bool, error) {
	c, err := database.Client.ThreadKudo.
		Query().
		Where(threadkudo.ThreadID(id)).
		Where(threadkudo.UserID(userId)).
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
	q := chi.URLParam(r, "threadId")
	threadId := 0

	if len(q) != 0 {
		threadId, _ = strconv.Atoi(q)
	}

	threads, err := GetThreads(ctx, threadId)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrRetrieveThreads, ReadHandler)
			res.Message = WrapErrRetrieveThreads.Message
			return res, err
		}
	
	userId, _ := users.GetUserIDFromToken(r)

	for _, thread := range threads {
		t, err := GetTagsByThread(ctx, thread)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrRetrieveTags, ReadHandler)
			res.Message = WrapErrRetrieveTags.Message
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
			res.Error = api.BuildError(err, WrapErrGetKudoCount, ReadHandler)
			res.Message = WrapErrGetKudoCount.Message
			return res, err
		}

		b, err := CheckDidUserKudo(ctx, thread.ID, userId)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrCheckDidUserKudo, ReadHandler)
			res.Message = WrapErrCheckDidUserKudo.Message
			return res, err
		}

		u, err := users.GetUsernameFromID(ctx, database.Client, thread.CreatedBy)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrGetUsernameFromID, ReadHandler)
			res.Message = WrapErrGetUsernameFromID.Message
			return res, err
		}

		topic, err := topics.ListByThread(ctx, thread)
		if err != nil {
			res.Error = api.BuildError(err, WrapErrGetTopicFromThread, ReadHandler)
			res.Message = WrapErrGetTopicFromThread.Message
			return res, err
		}

		data = append(data, ThreadsResponse{
			ID:          thread.ID,
			Title:       thread.Title,
			Slug:        thread.Slug,
			Description: thread.Description,
			ModifiedAt:  thread.ModifiedAt,
			CreatedByID:   thread.CreatedBy,
			CreatedByUsername: u,
			CreatedAt:   thread.CreatedAt,
			Tags:        &formattedTags,
			KudoCount:   c,
			UserKudoed:  b,
			TopicSlug: topic.Slug,
		})
	}

	encodedData, err := json.Marshal(data)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrEncodeView, ReadHandler)
		res.Message = WrapErrEncodeView.Message
	} else {
		res.Data = encodedData
		res.Message = SuccessfulListThreadsMessage
	}

	return res, err
}
