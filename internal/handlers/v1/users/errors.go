package users

import (
	"errors"

	"github.com/interngowhere/web-backend/internal/api"
)

// Custom errors
var (
	ErrUsernameAlreadyRegistered = errors.New("username has already been taken up by someone else")
	ErrEmailAlreadyRegistered    = errors.New("email has already been used to register for another account")
	ErrMissingInputField         = errors.New("there is one or more more missing input field(s) in the request body")
)

// Custom error message wrappers
var (
	WrapErrRequestFormat     = api.ErrorMessage{Message: "Invalid request format", Code: 400}
	WrapErrCreateUser        = api.ErrorMessage{Message: "Failed to create user", Code: 500}
	WrapErrCheckEmail        = api.ErrorMessage{Message: "Failed to check if email has been used to register for another account", Code: 500}
	WrapErrCheckUsername     = api.ErrorMessage{Message: "Failed to check if username has already been taken up by someone else", Code: 500}
	WrapErrHashPassword      = api.ErrorMessage{Message: "Failed to hash given password", Code: 500}
	WrapErrGetUser           = api.ErrorMessage{Message: "Failed to retrieve user from given email", Code: 500}
	WrapErrIncorrectPassword = api.ErrorMessage{Message: "Password provided is incorrect", Code: 400}
	WrapErrSetExpiry         = api.ErrorMessage{Message: "Failed to set JWT expiration", Code: 500}
	WrapErrEncodeJWT         = api.ErrorMessage{Message: "Failed to encode JWT", Code: 500}
	WrapErrEncodeView        = api.ErrorMessage{Message: "Failed to obtain JSON encoding of API response", Code: 500}
	WrapErrRetrieveUserID    = api.ErrorMessage{Message: "Failed to retrieve user ID from JWT", Code: 500}
	WrapErrDecodeRequest     = api.ErrorMessage{Message: "Failed to decode request body", Code: 500}
)
