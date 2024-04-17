package repository_test

import (
	"database/sql"
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
		expectedSqlErr    error
	}{
		{
			name:              "create token",
			expectedToken:     "token.token.token",
			expectedUserEmail: "mock@mock.com",
			expectedSqlErr:    sql.ErrNoRows,
		},
	}

	for _, test := range testCases {
		query := `INSERT INTO token (Token, UserEmail) VALUES (?, ?)`
		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		mock.
			ExpectExec(query).
			WithArgs(test.expectedToken, test.expectedUserEmail).
			WillReturnError(test.expectedSqlErr).
			WillReturnResult(sqlmock.NewResult(1, 1))
		defer db.Close()

		repo := repository.NewAuthRepository(db)
		_, createErr := repo.CreateToken(test.expectedToken, test.expectedUserEmail)
		sqlErr := mock.ExpectationsWereMet()

		assert.AssertError(t, sqlErr, nil)
		assert.AssertError(t, createErr, test.expectedSqlErr)
	}
}
