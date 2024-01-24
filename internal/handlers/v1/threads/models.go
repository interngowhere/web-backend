package threads

import (
	"time"

	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/internal/handlers/v1/tags"
)

// ThreadRequest is derived from ent.Thread but with
// modified json keys in camelCase to match JSON convention.
type ThreadRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        []int  `json:"tags"`
}

// ThreadRequest is derived from ent.Thread but with
// modified json keys in camelCase to match JSON convention.
// It also contains additional fields Tags, KudoCount, and UserKudoed
type ThreadsResponse struct {
	ID          int                 `json:"id"`
	Title       string              `json:"title,omitempty"`
	Slug        string              `json:"slug,omitempty"`
	Description string              `json:"description,omitempty"`
	ModifiedAt  time.Time           `json:"modifiedAt,omitempty"`
	CreatedByID   uuid.UUID           `json:"createdByID,omitempty"`
	CreatedByUsername string        `json:"createdByUsername,omitempty"`
	CreatedAt   time.Time           `json:"createdAt,omitempty"`
	Tags        *[]tags.TagResponse `json:"tags,omitempty"`
	KudoCount   int                 `json:"kudoCount,omitempty"`
	UserKudoed  bool                `json:"userKudoed,omitempty"`
	TopicSlug 	 string              `json:"topicSlug,omitempty"`
}
