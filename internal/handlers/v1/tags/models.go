package tags

type TagRequest struct {
	TagName string `json:"tagName"`
}

type TagResponse struct {
	ID      int    `json:"id,omitempty"`
	TagName string `json:"tagName,omitempty"`
}
