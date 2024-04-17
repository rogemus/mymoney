package repository_test

import (
	"database/sql"
	"testing"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_UserRepo_GetUser(t *testing.T) {
	var empty_user model.User
	user := mocks.GenerateUser(1)

	testCases := []struct {
		name           string
		expected       model.User
		expectedSqlErr error
		userId         int
	}{
		{
			name:           "returns user object",
			expected:       user,
			expectedSqlErr: nil,
			userId:         1,
		},
		{
			name:           "return 404 if user not found",
			expected:       empty_user,
			expectedSqlErr: sql.ErrNoRows,
			userId:         9999,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			query := "SELECT ID, Uuid, Email, Password, Username, Created FROM user WHERE ID = ?"
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
				WithArgs(test.userId).
				WillReturnRows(expectedRow).
				WillReturnError(test.expectedSqlErr)
			defer db.Close()

			userRepo := repository.NewUserRepository(db)
			user, getErr := userRepo.GetUser(test.userId)
			sqlErr := mock.ExpectationsWereMet()

			assert.AssertError(t, getErr, nil)
			assert.AssertError(t, sqlErr, nil)
			assert.AssertStruct[model.User](t, user, test.expected)
		})
	}
}
