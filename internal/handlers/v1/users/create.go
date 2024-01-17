package users

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/ent/user"
	"github.com/interngowhere/web-backend/internal/api"
	"github.com/interngowhere/web-backend/internal/database"
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

	r.ParseForm()
	if len(r.Form["email"]) == 0 ||
		len(r.Form["password"]) == 0 ||
		len(r.Form["username"]) == 0 {
		err := fmt.Errorf("there is one or more more missing form input(s) in the request. Received: %v", r.Form)
		res.Error = api.BuildError(err, WrapErrRequestFormat, CreateHandler)
		res.Message = WrapErrRequestFormat.Message
		return res, err
	}

	email := r.Form["email"][0]
	password := r.Form["password"][0]
	username := r.Form["username"][0]

	// check if email has already been registerd
	c, err := database.Client.User.
		Query().
		Where(
			user.
				EmailEQ(email),
		).
		Count(ctx)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrCheckEmail, CreateHandler)
		res.Message = WrapErrCheckEmail.Message
		return res, err
	}
	if c != 0 {
		e := WrapErrCreateUser
		e.Code = 400
		res.Error = api.BuildError(ErrEmailAlreadyRegistered, e, CreateHandler)
		res.Message = ErrEmailAlreadyRegistered.Error()
		return res, ErrUsernameAlreadyRegistered
	}

	// check if username has already been registered
	c, err = database.Client.User.
		Query().
		Where(
			user.
				UsernameEQ(username),
		).
		Count(ctx)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrCheckUsername, CreateHandler)
		res.Message = WrapErrCheckUsername.Message
		return res, err
	}
	if c != 0 {
		e := WrapErrCreateUser
		e.Code = 400
		res.Error = api.BuildError(ErrUsernameAlreadyRegistered, e, CreateHandler)
		res.Message = ErrUsernameAlreadyRegistered.Error()
		return res, ErrUsernameAlreadyRegistered
	}

	// hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrHashPassword, CreateHandler)
		res.Message = WrapErrHashPassword.Message
		return res, err
	}

	u := ent.User{
		Email:    email,
		Username: username,
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
