package users

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/user"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
)

const CreateHandler = "users.HandleCreate"
const SuccessfulCreateUserMessage = "User created successfully"

func CreateUser(ctx context.Context, client *ent.Client, user ent.User) error {
	err := client.User.
		Create().
		SetEmail(user.Email).
		SetHash(user.Hash).
		SetUsername(user.Username).
		SetCreatedAt(time.Now()).
		Exec(ctx)
	return err
}

func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	// Read JSON body in request into a new UserRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data UserRequest
	err := decoder.Decode(&data)
	// if err != nil {
	// 	res = api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
	// 	res.Message = customerrors.WrapErrDecodeRequest.Message
	// 	return res, err
	// }
	if err != nil {
		res := api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
		return res, err
	}

	// Check for missing input fields in request body
	if len(data.Email) == 0 ||
		len(data.Password) == 0 ||
		len(data.Username) == 0 {
		res = api.BuildError(customerrors.ErrMalformedRequest, customerrors.WrapErrRequestFormat, CreateHandler)
		return res, customerrors.ErrMalformedRequest
	}

	// check if email has already been registerd
	c, err := database.Client.User.
		Query().
		Where(
			user.
				EmailEQ(data.Email),
		).
		Count(ctx)
	if err != nil {
		res = api.BuildError(err, WrapErrCheckEmail, CreateHandler)
		return res, err
	}
	if c != 0 {
		res = api.BuildError(ErrEmailAlreadyRegistered, customerrors.WrapErrResourceExist, CreateHandler)
		return res, ErrEmailAlreadyRegistered
	}

	// check if username has already been registered
	c, err = database.Client.User.
		Query().
		Where(
			user.
				UsernameEQ(data.Username),
		).
		Count(ctx)
	if err != nil {
		res = api.BuildError(err, WrapErrCheckUsername, CreateHandler)
		return res, err
	}
	if c != 0 {
		res = api.BuildError(ErrUsernameAlreadyRegistered, customerrors.WrapErrResourceExist, CreateHandler)
		return res, ErrUsernameAlreadyRegistered
	}

	// hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		res = api.BuildError(err, WrapErrHashPassword, CreateHandler)
		return res, err
	}

	u := ent.User{
		Email:    data.Email,
		Username: data.Username,
		Hash:     hashedPassword,
	}

	err = CreateUser(ctx, database.Client, u)
	if err != nil {
		res = api.BuildError(err, WrapErrCreateUser, CreateHandler)
		return res, err
	}

	res.Message = SuccessfulCreateUserMessage
	res.Code = 201

	return res, err
}
