package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "Get All Users",
			route:         "/backend-go/v1/user",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"data\":[]}",
		},
	}

	app := SetupFiberApp()

	for _, test := range tests {
		req, _ := http.NewRequest("GET", test.route, nil)
		res, err := app.Test(req, -1)
		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
		body, err := io.ReadAll(res.Body)
		assert.Nilf(t, err, test.description)
		assert.Equal(t, test.expectedBody, string(body), test.description)
	}
}
