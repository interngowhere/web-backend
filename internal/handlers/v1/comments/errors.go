package comments

import (
	"errors"

	"github.com/interngowhere/web-backend/internal/api"
)

// Custom errors
var (
	ErrNoMatchFromThreadID  = errors.New("no comment is found associated to the thread id provided")
	ErrNoMatchFromCommentID = errors.New("no comment is found associated to the comment id provided")
	ErrCommentKudoExist     = errors.New("the user has already given a kudo to the same comment")
)

// Custom error message wrappers
var (
	WrapErrCheckCommentKudo = api.ErrorMessage{Message: "An error occured while checking if the user has already given a kudo to the same comment", Code: 500}
	WrapErrCommentKudoExist = api.ErrorMessage{Message: "the user has already given a kudo to the same comment", Code: 400}

	WrapErrRetrieveComments = api.ErrorMessage{Message: "Failed to retrieve comment(s)", Code: 500}
	WrapErrCreateComment    = api.ErrorMessage{Message: "Failed to create comment", Code: 500}
	WrapErrUpdateComment    = api.ErrorMessage{Message: "Failed to update comment", Code: 500}
	WrapErrDeleteComment    = api.ErrorMessage{Message: "Failed to delete comment", Code: 500}
	WrapErrAddKudo          = api.ErrorMessage{Message: "Failed to add kudo", Code: 500}
	WrapErrRemoveKudo       = api.ErrorMessage{Message: "Failed to remove kudo", Code: 500}

	WrapErrGetKudoCount      = api.ErrorMessage{Message: "Failed to get kudo count on the comment", Code: 500}
	WrapErrCheckDidUserKudo  = api.ErrorMessage{Message: "Something went wrong while checking if user gave a kudo to the comment", Code: 500}
	WrapErrGetUserFromID = api.ErrorMessage{Message: "Failed to get username from ID", Code: 500}
)
