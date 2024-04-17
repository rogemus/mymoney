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
		user           model.User
	}{
		{
			name:           "create budget",
			expectedErr:    nil,
			expectedSqlErr: nil,
			user: model.User{
				Username: "test",
				Password: "pass",
				Email:    "test@test.com",
				ID:       1,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer db.Close()

			mock.
				ExpectExec("INSERT INTO user (Username, Email, Password) VALUES (?, ?, ?)").
				WithArgs(test.user.Username, test.user.Email, test.user.Password).
				WillReturnResult(sqlmock.NewResult(int64(test.user.ID), 1)).
				WillReturnError(test.expectedSqlErr)
			defer db.Close()

			repo := repository.NewUserRepository(db)
			newUserId, createUsrErr := repo.CreateUser(test.user)
			sqlErr := mock.ExpectationsWereMet()

			assert.AssertError(t, createUsrErr, nil)
			assert.AssertError(t, sqlErr, nil)
			assert.AssertInt(t, int(newUserId), test.user.ID)
		})
	}
}
