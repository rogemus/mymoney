package handlers_test

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"tracker/pkg/handler"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockToken struct{}

func (m MockToken) Match(v driver.Value) bool {
	ok, _ := regexp.Match(`\S+\.{1}\S+\.{1}\S+`, []byte(v.(string)))
	return ok
}

func Test_UserHandler_Login(t *testing.T) {
	user := mocks.GenerateUser(1)
	token := MockToken{}
	var empty_user model.User

	testCases := []struct {
		name           string
		expected       string
		expectedStatus int
		body           string
		userEmail      string
		userPass       string
		token          MockToken
		user           model.User
	}{
		{
			name:           "returns token after login",
			expected:       `{"token":"token.token.token"}`,
			expectedStatus: 200,
			body:           `{"email": "mock@mock.com", "password": "pass"}`,
			userEmail:      "mock@mock.com",
			userPass:       "pass",
			token:          token,
			user:           user,
		},
		{
			name:           "returns error if empty pass",
			expected:       `{"msg":"Invalid or missing request params"}`,
			expectedStatus: 422,
			body:           `{"email": "mock@mock.com"}`,
			user:           empty_user,
		},
		{
			name:           "returns error if empty email",
			expected:       `{"msg":"Invalid or missing request params"}`,
			expectedStatus: 422,
			body:           `{"password": "pass123"}`,
			user:           empty_user,
		},
		{
			name:           "returns error if broken json",
			expected:       `{"msg":"Invalid request"}`,
			expectedStatus: 400,
			body:           `{broken`,
			user:           empty_user,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()

			columns := []string{
				"ID",
				"Uuid",
				"Email",
				"Password",
				"Username",
				"Created",
			}

			expectedRow := sqlmock.NewRows(columns)
			expectedRow.AddRow(
				&test.user.ID,
				&test.user.Uuid,
				&test.user.Email,
				&test.user.Password,
				&test.user.Username,
				&test.user.Created,
			)

			mock.
				ExpectQuery("SELECT").
				WithArgs(test.userEmail).
				WillReturnRows(expectedRow)

			mock.
				ExpectExec("INSERT").
				WithArgs(test.token, test.userEmail).
				WillReturnResult(sqlmock.NewResult(1, 1))

			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(test.body)))

			userRepo := repository.NewUserRepository(db)
			authRepo := repository.NewAuthRepository(db)

			userHandler := handler.NewUserHandler(userRepo, authRepo)

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(userHandler.LoginUser)
			hr.ServeHTTP(rr, req)

			if test.user.Email != "" {
				var token model.Token
				json.Unmarshal([]byte(rr.Body.String()), &token)
				assert.AssertRegex(t, token.Token, `\S+\.{1}\S+\.{1}\S+`)
			} else {
				assert.AssertJson(t, rr.Body.String(), test.expected)
			}

			assert.AssertInt(t, rr.Code, test.expectedStatus)
		})
	}
}
