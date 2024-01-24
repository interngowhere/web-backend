package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

type Response struct {
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   string          `json:"error,omitempty"`
	Code    int             `json:"code"`
}

// ErrorMessage encodes data for custom errors
// Example Usage:
// var ErrEncodeView = api.ErrorMessage{ Message: "Failed to retrieve topics", Code: 500 }
type ErrorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func BuildError(e error, m ErrorMessage, handler string) *Response {
	e = errors.Wrap(e, fmt.Sprintf("%s via %s", m.Message, handler))
	return &Response{
		Message: m.Message,
		Error:   e.Error(),
		Code:    m.Code,
	}
}

func BuildRouteHandler(f func(w http.ResponseWriter, r *http.Request) (*Response, error)) func(w http.ResponseWriter, req *http.Request) {
	g := func(w http.ResponseWriter, req *http.Request) {
		res, _ := f(w, req)
		if res.Code != 0 {
			render.Status(req, res.Code)
		} else {
			res.Code = 200
		}
		render.JSON(w, req, res)
	}
	return g
}
