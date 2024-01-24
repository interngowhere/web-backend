package api

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildError(t *testing.T) {
	e := errors.New("ERROR")
	m := ErrorMessage{
		Message: "Something went wrong",
		Code:    500,
	}
	h := "resource.HandleAction"

	require.Equal(
		t, 
		&Response{
			Message: "Something went wrong",
			Error:   "Something went wrong via resource.HandleAction: ERROR",
			Code:    500,
		}, 
		BuildError(e, m, h))
}

func TestBuildRouteHandler(t *testing.T) {
	f := func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		res := &Response{
			Message: "Something went wrong",
			Error:   "Something went wrong via resource.HandleAction: ERROR",
			Code:    500,
		}
		return res, nil;
	}
	require.HTTPBodyContains(t, BuildRouteHandler(f), "GET", "", nil, "Something went wrong via resource.HandleAction: ERROR")
	require.HTTPStatusCode(t, BuildRouteHandler(f), "GET", "", nil, 500)
}