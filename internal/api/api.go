package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

type ResponseError struct {
	Cause string `json:"cause"`
	Code  int    `json:"code"`
}

type Response struct {
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   *ResponseError  `json:"error,omitempty"`
}

// ErrorMessage encodes data for custom errors
// Example Usage:
// var ErrEncodeView = api.ErrorMessage{ Message: "Failed to retrieve topics", Code: 500 }
type ErrorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func BuildError(e error, m ErrorMessage, handler string) *ResponseError {
	e = errors.Wrap(e, fmt.Sprintf("%s via %s", m.Message, handler))
	return &ResponseError{
		Code:  m.Code,
		Cause: e.Error(),
	}
}

func BuildRouteHandler(f func(w http.ResponseWriter, r *http.Request) (*Response, error)) func(w http.ResponseWriter, req *http.Request) {
	g := func(w http.ResponseWriter, req *http.Request) {
		res, err := f(w, req)
		if err != nil {
			render.Status(req, res.Error.Code)
		}
		render.JSON(w, req, res)
	}
	return g
}
