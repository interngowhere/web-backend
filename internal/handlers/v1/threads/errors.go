package threads

import (
	"errors"

	"github.com/interngowhere/web-backend/internal/api"
)

// Custom errors
var (
	ErrThreadKudoExist = errors.New("the user has already given a kudo to the same thread")
	ErrNoThreadFound   = errors.New("no thread found")
)

// Custom error messages
var (
	WrapErrRetrieveTopic   = api.ErrorMessage{Message: "Failed to retrieve topic", Code: 500}
	WrapErrCheckTopicExist = api.ErrorMessage{Message: "An error occured while checking if the topic already exists", Code: 500}
	WrapErrCheckThreadKudo = api.ErrorMessage{Message: "An error occured while checking if the user has already given a kudo to the same thread", Code: 500}

	WrapErrThreadKudoExist = api.ErrorMessage{Message: "the user has already given a kudo to the same thread", Code: 400}

	WrapErrRetrieveThreads = api.ErrorMessage{Message: "Failed to retrieve thread(s)", Code: 500}
	WrapErrCreateThread    = api.ErrorMessage{Message: "Failed to create thread", Code: 500}
	WrapErrUpdateThread    = api.ErrorMessage{Message: "Failed to update thread", Code: 500}
	WrapErrDeleteThread    = api.ErrorMessage{Message: "Failed to delete thread", Code: 500}
	WrapErrAddKudo         = api.ErrorMessage{Message: "Failed to add kudo", Code: 500}
	WrapErrRemoveKudo      = api.ErrorMessage{Message: "Failed to remove kudo", Code: 500}

	WrapErrRetrieveTags       = api.ErrorMessage{Message: "Failed to get the tags associated with the thread", Code: 500}
	WrapErrGetKudoCount       = api.ErrorMessage{Message: "Failed to get kudo count on the thread", Code: 500}
	WrapErrCheckDidUserKudo   = api.ErrorMessage{Message: "Something went wrong while checking if user gave a kudo to the thread", Code: 500}
	WrapErrGetTopicFromThread = api.ErrorMessage{Message: "Failed to get topic from thread", Code: 500}
)
