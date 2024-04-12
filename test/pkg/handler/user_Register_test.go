package handlers_test

import (
	"bytes"
	"database/sql/driver"
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

type MockPass struct{}

func (m MockPass) Match(v driver.Value) bool {
	ok, _ := regexp.Match(`\S+`, []byte(v.(string)))
	return ok
}

func Test_UserHandler_Register(t *testing.T) {
	user := mocks.GenerateUser(1)
  mock_pass := MockPass{}
	var empty_user model.User

	testCases := []struct {
		name           string
		expected       string
		expectedStatus int
		user           model.User
		body           string
	}{
		{
			name:           "returns msg after create",
			expected:       `{"msg":"User created"}`,
			expectedStatus: 201,
			body:           `{"username": "Mock Mosinski", "email": "mock@mock.com", "password": "pass"}`,
			user:           user,
		},
		{
			name:           "returns error if empty email",
			expected:       `{"msg":"Invalid or missing request params"}`,
			expectedStatus: 422,
			body:           `{"username": "Mock Mosinski", "password": "pass"}`,
			user:           empty_user,
		},
		{
			name:           "returns error if empty password",
			expected:       `{"msg":"Invalid or missing request params"}`,
			expectedStatus: 422,
			body:           `{"username": "Mock Mosinski", "email": "mock@mock.com"}`,
			user:           empty_user,
		},
		{
			name:           "returns error if empty username",
			expected:       `{"msg":"Invalid or missing request params"}`,
			expectedStatus: 422,
			body:           `{"email": "mock@mock.com", "password": "pass"}`,
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

			mock.
				ExpectExec("INSERT").
				WithArgs(test.user.Username, test.user.Email, mock_pass).
				WillReturnResult(sqlmock.NewResult(1, 1))

			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(test.body)))

			userRepo := repository.NewUserRepository(db)
			authRepo := repository.NewAuthRepository(db)

			userHandler := handler.NewUserHandler(userRepo, authRepo)

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(userHandler.RegisterUser)
			hr.ServeHTTP(rr, req)

			assert.AssertJson(t, rr.Body.String(), test.expected)
			assert.AssertInt(t, rr.Code, test.expectedStatus)
		})
	}
}
