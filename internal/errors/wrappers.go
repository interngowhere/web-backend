package customerrors

import "github.com/interngowhere/web-backend/internal/api"

// Custom error message wrappers
var (
	// 4xx
	WrapErrRequestFormat = api.ErrorMessage{Message: "Invalid request format", Code: 400}
	WrapErrNotFound      = api.ErrorMessage{Message: "Resource cannot be found", Code: 404}
	WrapErrResourceExist = api.ErrorMessage{Message: "Resource already exists", Code: 409}

	// 5xx
	WrapErrDecodeRequest = api.ErrorMessage{Message: "Failed to decode JSON request body", Code: 500}
	WrapErrEncodeView    = api.ErrorMessage{Message: "Failed to encode API response to JSON", Code: 500}
	WrapErrStrToInt      = api.ErrorMessage{Message: "Failed to convert type string to type int", Code: 500}
)
