package users

import "time"

// UserRequest models the JSON request body
// for creating or updating a new user.
type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// LoginRequest models the JSON request body
// for logging in a user.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserResponse is derived from ent.User
// with modified json keys in camelCase.
type CreateUserResponse struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Slug             string    `json:"slug"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"shortDescription"`
	ProfilePicURL    string    `json:"profilePicURL"`
	CreatedAt        time.Time `json:"createdAt"`
}
