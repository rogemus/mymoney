package repository_test

import (
	"testing"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_AuthRepo_CreateToken(t *testing.T) {
	testCases := []struct {
		name              string
		expectedToken     string
		expectedUserEmail string
		expectedErr       error
		expectedSqlErr    error
	}{
		{
			name:              "create token",
			expectedToken:     "token.token.token",
			expectedUserEmail: "mock@mock.com",
			expectedErr:       nil,
			expectedSqlErr:    nil,
		},
	}

	for _, test := range testCases {
		query := `INSERT INTO tokens (token, useremail) VALUES ($1, $2)`
		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		mock.
			ExpectExec(query).
			WithArgs(test.expectedToken, test.expectedUserEmail).
			WillReturnError(test.expectedSqlErr).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := repository.NewAuthRepository(db)
		createErr := repo.CreateToken(test.expectedToken, test.expectedUserEmail)

		sqlErr := mock.ExpectationsWereMet()
		assert.AssertError(t, sqlErr, nil)
		assert.AssertError(t, createErr, test.expectedErr)
	}
}
