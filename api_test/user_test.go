package api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/backend-go/internal/app"
	"github.com/backend-go/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	var actualResp struct {
		Data []user.User `json:"data"`
	}
	var expectedUsers []user.User

	expectedBytes, err := os.ReadFile("../resources/users.json")
	if err != nil {
		t.Fatalf("cannot read users.json: %v", err)
	}

	app := app.SetupFiberApp()

	expectedBody := strings.TrimSpace(string(expectedBytes))

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
			expectedBody:  expectedBody,
		},
	}

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

		if err := json.Unmarshal(body, &actualResp); err != nil {
			t.Fatalf("cannot unmarshal response body: %v", err)
		}

		if err := json.Unmarshal([]byte(test.expectedBody), &expectedUsers); err != nil {
			t.Fatalf("cannot unmarshal users.json: %v", err)
		}
		assert.Equal(t, len(expectedUsers), len(actualResp.Data), test.description)

		for i, u := range expectedUsers {
			assert.Equal(t, u.Id, actualResp.Data[i].Id, test.description+": user Id should match")
			assert.Equal(t, u.FirstName, actualResp.Data[i].FirstName, test.description+": FirstName should match")
			assert.Equal(t, u.LastName, actualResp.Data[i].LastName, test.description+": LastName should match")
			assert.Equal(t, u.Username, actualResp.Data[i].Username, test.description+": Username should match")
			assert.Equal(t, u.Password, actualResp.Data[i].Password, test.description+": Password should match")
		}
	}
}
