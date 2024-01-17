package tags

import (
	"errors"

	"github.com/interngowhere/web-backend/internal/api"
)

// Custom errors
var (
	ErrMissingTagName = errors.New("tag name is required but not provided in the request body")
	ErrTagExist = errors.New("the tag already exists")
)

// Custom error messages
var (
	WrapErrRequestFormat   = api.ErrorMessage{Message: "Invalid request format", Code: 400}
	WrapErrEncodeView      = api.ErrorMessage{Message: "Failed to encode API response in JSON", Code: 500}

	WrapErrCheckTagExist = api.ErrorMessage{Message: "Failed to check if tag already exists", Code: 500}

	WrapErrNoTagFound	= api.ErrorMessage{Message: "No tag found", Code: 404}
	WrapErrRetrieveTags  = api.ErrorMessage{Message: "Failed to retrieve tag(s)", Code: 500}
	WrapErrCreateTag     = api.ErrorMessage{Message: "Failed to create tag", Code: 500}
	WrapErrUpdateTag     = api.ErrorMessage{Message: "Failed to update tag", Code: 500}
	WrapErrDeleteTag     = api.ErrorMessage{Message: "Failed to delete tag", Code: 500}
	
	WrapErrStrToInt 		= api.ErrorMessage{Message: "Failed to convert type string to type int", Code: 500}
	WrapErrDecodeRequest = api.ErrorMessage{Message: "Failed to decode request body", Code: 500}
	WrapErrRetrieveUserID   = api.ErrorMessage{Message: "Failed to retrieve user ID from JWT", Code: 500}
)