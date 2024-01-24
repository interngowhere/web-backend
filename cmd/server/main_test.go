package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-chi/jwtauth/v5"
	jwt "github.com/interngowhere/web-backend/internal/auth"
	"github.com/interngowhere/web-backend/internal/database"
	customerrors "github.com/interngowhere/web-backend/internal/errors"
	"github.com/interngowhere/web-backend/internal/handlers/v1/topics"
	"github.com/interngowhere/web-backend/internal/middleware"
	"github.com/interngowhere/web-backend/internal/router"
	"github.com/interngowhere/web-backend/internal/server"
	"github.com/stretchr/testify/require"
)

var s *server.Server = nil

// executeRequest creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
// SOURCE: go-chi.io
func executeRequest(req *http.Request, s *server.Server) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    s.Router.ServeHTTP(rr, req)

    return rr
}

// assertStatusCode checks the status code of the response
// SOURCE: go-chi.io
func assertStatusCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func GenerateTestJWT() (string, error) {
	claims := map[string]interface{}{"id": "123"}
	err := jwt.SetExpiry(claims)
	if err != nil {
		return "", fmt.Errorf("error setting JWT expiry: %w", err)
	}
	_, tokenString, err := jwt.TokenAuth.Encode(claims)
	if err != nil {
		return "", fmt.Errorf("error encoding JWT: %w", err)
	}
    return "Bearer " + tokenString, nil
}

func TestMain(t *testing.T) {
    // Set up server
	s = server.New()
	r := s.Router
	jwt.Init(os.Getenv("JWT_SIGKEY"))
	middleware.Setup(r)
	router.Setup(r)

	// Set up database
	database.Init()
	defer database.Client.Close()

	// Test public route
    req, _ := http.NewRequest("GET", "/ping", nil)
    response := executeRequest(req, s)
    assertStatusCode(t, http.StatusOK, response.Code)
    require.Equal(t, ".", response.Body.String())

	// Test protected route without token
	req, _ = http.NewRequest("POST", "/ping", nil)
    response = executeRequest(req, s)
    assertStatusCode(t, http.StatusUnauthorized, response.Code)
    require.Equal(t, "no token found\n", response.Body.String())

	// Test protected route with token
	token, err := GenerateTestJWT()
	if err != nil {
		t.Errorf("error generating test JWT: %v", err)
	}

	req, _ = http.NewRequest("POST", "/ping", nil)
	req.Header.Add("Authorization", token)
    response = executeRequest(req, s)
    assertStatusCode(t, http.StatusOK, response.Code)
    require.Equal(t, ".", response.Body.String())

	// Test protected route with expired token
	claims := map[string]interface{}{"id": "123"}
	jwtauth.SetExpiry(claims, time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC))
	_, tokenString, _ := jwt.TokenAuth.Encode(claims)
    token = "Bearer " + tokenString

	req, _ = http.NewRequest("POST", "/ping", nil)
	req.Header.Add("Authorization", token)
    response = executeRequest(req, s)
    assertStatusCode(t, http.StatusUnauthorized, response.Code)
    require.Equal(t, "token is expired\n", response.Body.String())

}

func TestTopics(t *testing.T) {
	database.Init()
	defer database.Client.Close()
	
	// GET /topics with no result
    req, _ := http.NewRequest("GET", "/topics", nil)
    response := executeRequest(req, s)

    assertStatusCode(t, http.StatusNotFound, response.Code)
	m := customerrors.WrapErrNotFound.Message
	e := topics.ErrNoTopicFound.Error()
    require.Equal(
		t, 
		fmt.Sprintf(`{"message":"%s","error":"%s via topics.HandleList: %s","code":%s}` + "\n", m, m, e, strconv.Itoa(http.StatusNotFound)), 
		response.Body.String(),
	)

	// POST /topics with missing fields
	token, err := GenerateTestJWT()
	if err != nil {
		t.Errorf("error generating test JWT: %v", err)
	}

	var jsonStr = []byte(`{"title":"test"}`)
	req, _ = http.NewRequest("POST", "/topics", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", token)
    response = executeRequest(req, s)
    assertStatusCode(t, http.StatusBadRequest, response.Code)
	m = customerrors.WrapErrRequestFormat.Message
	e = customerrors.ErrMalformedRequest.Error()
    require.Equal(t, 
		fmt.Sprintf(`{"message":"%s","error":"%s via topics.HandleCreate: %s","code":%s}` + "\n", m, m, e, strconv.Itoa(http.StatusBadRequest)), 
		response.Body.String(),
	)

	// TODO: POST /topics with valid fields

	// TODO: POST /topics with conflict

	// TODO: GET /topics with result

	// TODO: GET /topics/{title} with no result

	// TODO: GET /topics/{title} with result

	// TODO: PUT /topics/{title}

	// TODO: PUT /topics/{title} with no match

	// TODO: DELETE /topics/{title}
}
