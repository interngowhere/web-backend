package comments

import (
	"time"

	"github.com/google/uuid"
)

type CommentRequest struct {
	ParentID string `json:"parentId"`
	Content  string `json:"content"`
}

type CommentsResponse struct {
	ID                int                 `json:"id"`
	ParentID          int                 `json:"parentId"`
	Content           string              `json:"content,omitempty"`
	ModifiedAt        time.Time           `json:"modifiedAt,omitempty"`
	CreatedByID       uuid.UUID           `json:"createdByID,omitempty"`
	CreatedByUsername string              `json:"createdByUsername,omitempty"`
	CreatedAt         time.Time           `json:"createdAt,omitempty"`
	KudoCount         int                 `json:"kudoCount"`
	UserKudoed        bool                `json:"userKudoed"`
	Children          *[]CommentsResponse `json:"children"`
}
