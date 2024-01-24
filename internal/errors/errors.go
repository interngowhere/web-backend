package customerrors

import "errors"

// Custom errors
var (
	ErrMalformedRequest = errors.New("there is one or more more missing/malformed field(s) in the request body")
	ErrResourceNotFound = errors.New("the requested resource was not found")
)
