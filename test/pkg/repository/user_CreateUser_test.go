package repository_test

import (
	"testing"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_UserRepo_CreateUser(t *testing.T) {
	testCases := []struct {
		name           string
		expectedErr    error
		expectedSqlErr error
		userName       string
		userPass       string
		userEmail      string
		userId         int
	}{
		{
			name:           "create budget",
			expectedErr:    nil,
			expectedSqlErr: nil,
			userName:       "test",
			userPass:       "pass",
			userEmail:      "test@test.com",
			userId:         1,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer db.Close()

			mock.
				ExpectExec("INSERT INTO user (Username, Email, Password) VALUES (?, ?, ?)").
				WithArgs(test.userName, test.userEmail, test.userPass).
				WillReturnResult(sqlmock.NewResult(int64(test.userId), 1)).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewUserRepository(db)

			newUser := model.User{Username: test.userName, Email: test.userEmail, Password: test.userPass}
			newUserId, createUsrErr := repo.CreateUser(newUser)
			err := mock.ExpectationsWereMet()

			assert.AssertInt(t, int(newUserId), test.userId)
			assert.AssertError(t, err, test.expectedSqlErr)
			assert.AssertError(t, createUsrErr, test.expectedErr)
		})
	}
}
