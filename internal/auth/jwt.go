package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

func Init(k string) {
	TokenAuth = jwtauth.New("HS256", []byte(k), nil)
}

func SetExpiry(claims map[string]interface{}) error {
	d, err := strconv.Atoi(os.Getenv("JWT_VALIDITY"))
	if err == nil {
		jwtauth.SetExpiry(claims, time.Now().Add(time.Minute*time.Duration(d)))
	}
	jwtauth.SetIssuedNow(claims)
	return err
}
