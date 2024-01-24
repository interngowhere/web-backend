package users

import (
	"errors"

	"github.com/interngowhere/web-backend/internal/api"
)

// Custom errors
var (
	ErrUsernameAlreadyRegistered = errors.New("username has already been taken up by someone else")
	ErrEmailAlreadyRegistered    = errors.New("email has already been used to register for another account")
)

// Custom error message wrappers
var (
	WrapErrCreateUser        = api.ErrorMessage{Message: "Failed to create user", Code: 500}
	WrapErrCheckEmail        = api.ErrorMessage{Message: "Something went wrong while checking if email has been used to register for another account", Code: 500}
	WrapErrCheckUsername     = api.ErrorMessage{Message: "Something went wrong while checking if username has already been taken up by someone else", Code: 500}
	WrapErrGetUserFromEmail  = api.ErrorMessage{Message: "Something went wrong while retrieving user from given email", Code: 500}
	WrapErrHashPassword      = api.ErrorMessage{Message: "Failed to hash given password", Code: 500}
	WrapErrGetUser           = api.ErrorMessage{Message: "Failed to retrieve user from given email", Code: 500}
	WrapErrIncorrectPassword = api.ErrorMessage{Message: "Password provided is incorrect", Code: 400}
	WrapErrSetExpiry         = api.ErrorMessage{Message: "Failed to set JWT expiration", Code: 500}
	WrapErrEncodeJWT         = api.ErrorMessage{Message: "Failed to encode JWT", Code: 500}
	WrapErrRetrieveIDFromJWT = api.ErrorMessage{Message: "Failed to retrieve user ID from JWT", Code: 500}
)
