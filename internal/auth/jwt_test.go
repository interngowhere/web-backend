package jwt

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJWT(t *testing.T) {
	Init("test")
	require.NotNil(t, TokenAuth)

	claims := map[string]interface{}{"id": "hello"}
	d, err := strconv.Atoi(os.Getenv("JWT_VALIDITY"))
	exp := time.Now().UTC().Add(time.Minute*time.Duration(d)).Format("1970-01-01 23:59:59")
	if err != nil {
		t.Error(err)
	}
	SetExpiry(claims)
	_, tokenString, _ := TokenAuth.Encode(claims)
	token, _ := TokenAuth.Decode(tokenString)
	id, _ := token.Get("id")
	require.Equal(t, "hello", id)
	require.Equal(t, exp, token.Expiration().Format("1970-01-01 23:59:59"))
}