package users

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/interngowhere/web-backend/internal/api"
	jwt "github.com/interngowhere/web-backend/internal/auth"
	"github.com/interngowhere/web-backend/internal/database"
)

const LoginHandler = "users.HandleLogin"
const SuccessfulUserLoginMessage = "User logged in successfully"

type JWTResponse struct {
	Token string `json:"token"`
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
		res.Error = api.BuildError(err, WrapErrDecodeRequest, CreateHandler)
		res.Message = WrapErrDecodeRequest.Message
		return res, err
	}

	// Check for missing input fields in request body
	if len(data.Email) == 0 ||
		len(data.Password) == 0 {
		res.Error = api.BuildError(ErrMissingInputField, WrapErrRequestFormat, CreateHandler)
		res.Message = WrapErrRequestFormat.Message
		return res, ErrMissingInputField
	}

	// get user from email
	u, err := GetUserFromEmail(ctx, database.Client, data.Email)
	if err != nil {
		e := WrapErrGetUser
		if err.Error() == "ent: user not found" {
			e.Code = 404
		}
		res.Error = api.BuildError(err, e, LoginHandler)
		res.Message = e.Message
		return res, err
	}

	err = bcrypt.CompareHashAndPassword(u.Hash, []byte(data.Password))
	if err != nil {
		res.Error = api.BuildError(err, WrapErrIncorrectPassword, LoginHandler)
		res.Message = WrapErrIncorrectPassword.Message
		return res, err
	}

	claims := map[string]interface{}{"id": u.ID}
	err = jwt.SetExpiry(claims)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrSetExpiry, LoginHandler)
		res.Message = WrapErrSetExpiry.Message
		return res, err
	}

	_, tokenString, err := jwt.TokenAuth.Encode(claims)
	if err != nil {
		res.Error = api.BuildError(err, WrapErrEncodeJWT, LoginHandler)
		res.Message = WrapErrEncodeJWT.Message
		return res, err
	}

	encodedData, err := json.Marshal(JWTResponse{Token: tokenString})
	if err != nil {
		res.Error = api.BuildError(err, WrapErrEncodeView, LoginHandler)
		res.Message = WrapErrEncodeView.Message
	} else {
		res.Data = encodedData
		res.Message = SuccessfulUserLoginMessage
	}

	return res, err
}
