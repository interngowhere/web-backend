package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
	"golang.org/x/crypto/bcrypt"
)

const UpdateHandler = "users.HandleUpdate"
const SuccessfulUpdateUserMessage = "Successfully updated user"

// UpdateUser updates an existing user entry in the database
// based on user ID provided via the ent.User object
func UpdateUser(ctx context.Context, client *ent.Client, user *ent.User) error {
	return client.User.
		UpdateOne(user).
		SetEmail(user.Email).
		SetUsername(user.Username).
		SetHash(user.Hash).
		Exec(ctx)
}

// HandleCreate parses the PUT request form data, calls
// UpdateUser if needed and returns a JSON encoded API response
func HandleUpdate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	userID, err := uuid.Parse(chi.URLParam(r, "userID"))
	if err != nil {
		res = api.BuildError(err, WrapErrParseUserID, DeleteHandler)
		return res, err
	}

	u, err := GetUserFromID(ctx, database.Client, userID)
	if err != nil {
		res = api.BuildError(err, WrapErrGetUser, DeleteHandler)
		return res, err
	}
	if u == nil {
		res = api.BuildError(customerrors.ErrResourceNotFound, customerrors.WrapErrNotFound, DeleteHandler)
		res.Message = customerrors.WrapErrNotFound.Message
		return res, customerrors.ErrResourceNotFound
	}

	// Read JSON body in request into a new UserRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data UserRequest
	err = decoder.Decode(&data)
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
		return res, err
	}

	// Update fields in user object only if a new value for the field is provided
	if len(data.Email) != 0 {
		u.Email = data.Email
	}
	if len(data.Username) != 0 {
		u.Username = data.Username
	}
	if len(data.Password) != 0 {
		// hash password using bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			res = api.BuildError(err, WrapErrHashPassword, CreateHandler)
			res.Message = WrapErrHashPassword.Message
			return res, err
		}
		u.Hash = hashedPassword
	}

	err = UpdateUser(ctx, database.Client, u)
	if err != nil {
		res = api.BuildError(err, WrapErrUpdateUser, UpdateHandler)
		return res, err
	}

	res.Message = SuccessfulUpdateUserMessage

	return res, err
}
