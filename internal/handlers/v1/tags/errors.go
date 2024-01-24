package tags

import (
	"errors"

	"github.com/interngowhere/web-backend/internal/api"
)

// Custom errors
var (
	ErrTagExist   = errors.New("the tag already exists")
	ErrNoTagFound = errors.New("no tag found")
)

// Custom error messages
var (
	WrapErrCheckTagExist = api.ErrorMessage{Message: "Something went wrong while checking if tag already exists", Code: 500}

	WrapErrRetrieveTags = api.ErrorMessage{Message: "Failed to retrieve tag(s)", Code: 500}
	WrapErrCreateTag    = api.ErrorMessage{Message: "Failed to create tag", Code: 500}
	WrapErrUpdateTag    = api.ErrorMessage{Message: "Failed to update tag", Code: 500}
	WrapErrDeleteTag    = api.ErrorMessage{Message: "Failed to delete tag", Code: 500}
)
