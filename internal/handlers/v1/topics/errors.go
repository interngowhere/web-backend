package topics

import (
	"errors"

	"github.com/interngowhere/web-backend/internal/api"
)

// Custom errors
var (
	ErrTopicExist = errors.New("topic already exists")
	ErrMissingIdentifier = errors.New("missing topic title identifier. Expected 'slug'")
	ErrMissingInputField = errors.New("there is one or more more missing input field(s) in the request body")
	ErrNoMatchFromSlug = errors.New("no match from the slug provided")
)

// Custom error message wrappers
var (
	WrapErrDecodeRequest = api.ErrorMessage{Message: "Failed to decode request body", Code: 500}
	WrapErrRequestFormat   = api.ErrorMessage{Message: "Invalid request format", Code: 400}
	WrapErrNoTopicFound	   = api.ErrorMessage{Message: "No topic found", Code: 404}
	WrapErrCheckTopicExist = api.ErrorMessage{Message: "Failed to check if topic already exists", Code: 500}
	WrapErrRetrieveTopics  = api.ErrorMessage{Message: "Failed to retrieve topic(s)", Code: 500}
	WrapErrCreateTopic     = api.ErrorMessage{Message: "Failed to create topic", Code: 500}
	WrapErrUpdateTopic     = api.ErrorMessage{Message: "Failed to update topic", Code: 500}
	WrapErrDeleteTopic     = api.ErrorMessage{Message: "Failed to delete topic", Code: 500}
	WrapErrEncodeView      = api.ErrorMessage{Message: "Failed to retrieve topic(s)", Code: 500}
)
