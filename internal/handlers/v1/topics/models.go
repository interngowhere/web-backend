package topics

// TopicRequest is derived from ent.Topic
// with modified json keys in camelCase.
type TopicRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	ProfilePicURL    string `json:"profilePicURL"`
}
