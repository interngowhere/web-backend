package topics

import (
	"errors"

	"github.com/interngowhere/web-backend/internal/api"
)

// Custom errors
var (
	ErrTopicExist   = errors.New("topic already exists")
	ErrNoTopicFound = errors.New("no topic found")
)

// Custom error message wrappers
var (
	WrapErrCheckTopicExist = api.ErrorMessage{Message: "Something went wrong while checking if topic already exists", Code: 500}
	WrapErrRetrieveTopics  = api.ErrorMessage{Message: "Failed to retrieve topic(s)", Code: 500}
	WrapErrCreateTopic     = api.ErrorMessage{Message: "Failed to create topic", Code: 500}
	WrapErrUpdateTopic     = api.ErrorMessage{Message: "Failed to update topic", Code: 500}
	WrapErrDeleteTopic     = api.ErrorMessage{Message: "Failed to delete topic", Code: 500}
)
