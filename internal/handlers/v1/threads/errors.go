package threads

import (
	"errors"

	"github.com/interngowhere/web-backend/internal/api"
)

// Custom errors
var (
	ErrMissingTitle = errors.New("thread title is required but not provided in the request")
	ErrNoMatchFromID = errors.New("no match from the thread ID provided")
	ErrThreadKudoExist = errors.New("the user has already given a kudo to the same thread")
)

// Custom error messages
var (
	WrapErrRequestFormat   = api.ErrorMessage{Message: "Invalid request format", Code: 400}
	WrapErrEncodeView      = api.ErrorMessage{Message: "Failed to retrieve topic(s)", Code: 500}

	WrapErrRetrieveTopic   = api.ErrorMessage{Message: "Failed to retrieve topic", Code: 500}
	WrapErrNoTopicFound	   = api.ErrorMessage{Message: "No topic found", Code: 404}
	WrapErrCheckTopicExist = api.ErrorMessage{Message: "An error occured while checking if the topic already exists", Code: 500}
	WrapErrCheckThreadKudo= api.ErrorMessage{Message: "An error occured while checking if the user has already given a kudo to the same thread", Code: 500}

	WrapErrThreadKudoExist = api.ErrorMessage{Message: "the user has already given a kudo to the same thread", Code: 400}

	WrapErrNoThreadFound	= api.ErrorMessage{Message: "No thread found", Code: 404}
	WrapErrRetrieveThreads  = api.ErrorMessage{Message: "Failed to retrieve thread(s)", Code: 500}
	WrapErrCreateThread     = api.ErrorMessage{Message: "Failed to create thread", Code: 500}
	WrapErrUpdateThread     = api.ErrorMessage{Message: "Failed to update thread", Code: 500}
	WrapErrDeleteThread     = api.ErrorMessage{Message: "Failed to delete thread", Code: 500}
	WrapErrAddKudo    = api.ErrorMessage{Message: "Failed to add kudo", Code: 500}
	WrapErrRemoveKudo     = api.ErrorMessage{Message: "Failed to remove kudo", Code: 500}

	WrapErrStrToInt 		= api.ErrorMessage{Message: "Failed to convert type string to type int", Code: 500}
	WrapErrDecodeRequest = api.ErrorMessage{Message: "Failed to decode request body", Code: 500}
	WrapErrRetrieveUserID   = api.ErrorMessage{Message: "Failed to retrieve user ID from JWT", Code: 500}
	WrapErrRetrieveTags     = api.ErrorMessage{Message: "Failed to get the tags associated with the thread", Code: 500}
	WrapErrGetKudoCount     = api.ErrorMessage{Message: "Failed to get kudo count on the thread", Code: 500}
	WrapErrCheckDidUserKudo = api.ErrorMessage{Message: "Failed to check if user gave a kudo to the thread", Code: 500}
)