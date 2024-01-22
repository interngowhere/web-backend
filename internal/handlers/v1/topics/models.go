package topics

import "time"

// TopicRequest is derived from ent.Topic
// with modified json keys in camelCase.
type TopicRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	ProfilePicURL    string `json:"profilePicURL"`
}

// TopicResponse is derived from ent.Topic
// with modified json keys in camelCase.
type TopicResponse struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Slug             string    `json:"slug"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"shortDescription"`
	ProfilePicURL    string    `json:"profilePicURL"`
	CreatedAt        time.Time `json:"createdAt"`
}
