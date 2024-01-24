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

	// Read JSON body in request into a new CreateUserRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data CreateUserRequest
	err := decoder.Decode(&data)
	if err != nil {
		res.Error = api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
		res.Message = customerrors.WrapErrDecodeRequest.Message
		return res, err
	}

	// Check for missing input fields in request body
	if len(data.Email) == 0 ||
		len(data.Password) == 0 ||
		len(data.Username) == 0 {
		res.Error = api.BuildError(customerrors.ErrMalformedRequest, customerrors.WrapErrRequestFormat, CreateHandler)
		res.Message = customerrors.WrapErrRequestFormat.Message
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
		res.Error = api.BuildError(err, WrapErrCheckEmail, CreateHandler)
		res.Message = WrapErrCheckEmail.Message
		return res, err
	}
	if c != 0 {
		res.Error = api.BuildError(ErrEmailAlreadyRegistered, customerrors.WrapErrResourceExist, CreateHandler)
		res.Message = ErrEmailAlreadyRegistered.Error()
		return res, ErrUsernameAlreadyRegistered
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
		res.Error = api.BuildError(err, WrapErrCheckUsername, CreateHandler)
		res.Message = WrapErrCheckUsername.Message
		return res, err
	}
	if c != 0 {
		res.Error = api.BuildError(ErrUsernameAlreadyRegistered, customerrors.WrapErrResourceExist, CreateHandler)
		res.Message = ErrUsernameAlreadyRegistered.Error()
		return res, ErrUsernameAlreadyRegistered
	}

	// hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrHashPassword, CreateHandler)
		res.Message = WrapErrHashPassword.Message
		return res, err
	}

	u := ent.User{
		Email:    data.Email,
		Username: data.Username,
		Hash:     hashedPassword,
	}

	err = CreateUser(ctx, database.Client, u)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrCreateUser, CreateHandler)
		res.Message = WrapErrCreateUser.Message
		return res, err
	}

	res.Message = SuccessfulCreateUserMessage

	return res, err
}
