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

	r.ParseForm()
	email := r.Form["email"][0]
	password := r.Form["password"][0]

	// get user from email
	u, err := GetUserFromEmail(ctx, database.Client, email)
	if err != nil {
		e := WrapErrGetUser
		if err.Error() == "ent: user not found" {
			e.Code = 404
		}
		res.Error = api.BuildError(err, e, LoginHandler)
		res.Message = e.Message
		return res, err
	}

	err = bcrypt.CompareHashAndPassword(u.Hash, []byte(password))
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

	data, err := json.Marshal(JWTResponse{Token: tokenString})
	if err != nil {
		res.Error = api.BuildError(err, WrapErrEncodeView, LoginHandler)
		res.Message = WrapErrEncodeView.Message
	} else {
		res.Data = data
		res.Message = SuccessfulUserLoginMessage
	}

	return res, err
}
