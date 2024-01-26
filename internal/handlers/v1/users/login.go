package users

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/internal/api"
	jwt "github.com/interngowhere/web-backend/internal/auth"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
)

const LoginHandler = "users.HandleLogin"
const SuccessfulUserLoginMessage = "User logged in successfully"

type JWTResponse struct {
	Token string `json:"token"`
	UserID string `json:"userId"`
	// other JWT related payload can be added here
}

func HandleLogin(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	ctx := context.Background()
	res := &api.Response{}

	// Read JSON body in request into a new LoginRequest object for use
	decoder := json.NewDecoder(r.Body)
	var data LoginRequest
	err := decoder.Decode(&data)
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrDecodeRequest, CreateHandler)
		return res, err
	}

	// Check for missing input fields in request body
	if len(data.Email) == 0 ||
		len(data.Password) == 0 {
		res = api.BuildError(customerrors.ErrMalformedRequest, customerrors.WrapErrRequestFormat, CreateHandler)
		return res, customerrors.ErrMalformedRequest
	}

	// get user from email
	u, err := GetUserFromEmail(ctx, database.Client, data.Email)
	if err != nil {
		switch t := err.(type) {
		default:
			res = api.BuildError(t, WrapErrGetUser, LoginHandler)
		case *ent.NotFoundError:
			res = api.BuildError(t, WrapErrInvalidEmail, LoginHandler)
		}
		return res, err
	}

	// Compare password hash
	err = bcrypt.CompareHashAndPassword(u.Hash, []byte(data.Password))
	if err != nil {
		res = api.BuildError(err, WrapErrIncorrectPassword, LoginHandler)
		return res, err
	}

	// Generate JWT token
	claims := map[string]interface{}{"id": u.ID}
	err = jwt.SetExpiry(claims)
	if err != nil {
		res = api.BuildError(err, WrapErrSetExpiry, LoginHandler)
		return res, err
	}

	_, tokenString, err := jwt.TokenAuth.Encode(claims)
	if err != nil {
		res = api.BuildError(err, WrapErrEncodeJWT, LoginHandler)
		return res, err
	}

	encodedData, err := json.Marshal(JWTResponse{Token: tokenString, UserID: u.ID.String()})
	if err != nil {
		res = api.BuildError(err, customerrors.WrapErrEncodeView, LoginHandler)
		res.Message = customerrors.WrapErrEncodeView.Message
	} else {
		res.Data = encodedData
		res.Message = SuccessfulUserLoginMessage
	}

	return res, err
}
