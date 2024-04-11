package repository_test

import (
	"database/sql"
	"testing"
	"tracker/pkg/errs"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_UserRepo_GetUserByEmail(t *testing.T) {
	var empty_user model.User
	user := mocks.GenerateUser(1)

	testCases := []struct {
		name           string
		expected       model.User
		expectedErr    error
		expectedSqlErr error
		userEmail      string
	}{
		{
			name:           "returns user object",
			expected:       user,
			expectedErr:    nil,
			expectedSqlErr: nil,
			userEmail:      "mock@mock.com",
		},
		{
			name:           "return 404 if user not found",
			expected:       empty_user,
			expectedErr:    errs.User404Err,
			expectedSqlErr: sql.ErrNoRows,
			userEmail:      "error@error.com",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			query := "SELECT ID, Uuid, Email, Password, Username, Created FROM user WHERE Email = ?"
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

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
				&test.expected.ID,
				&test.expected.Uuid,
				&test.expected.Email,
				&test.expected.Password,
				&test.expected.Username,
				&test.expected.Created,
			)

			mock.
				ExpectQuery(query).
				WithArgs(test.userEmail).
				WillReturnRows(expectedRow).
				WillReturnError(test.expectedSqlErr)

			userRepo := repository.NewUserRepository(db)
			user, getErr := userRepo.GetUserByEmail(test.userEmail)
			sqlErr := mock.ExpectationsWereMet()

			assert.AssertError(t, getErr, test.expectedErr)
			assert.AssertError(t, sqlErr, nil)
			assert.AssertStruct[model.User](t, user, test.expected)
		})
	}
}
