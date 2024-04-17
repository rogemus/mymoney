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

func Test_AuthRepo_GetToken(t *testing.T) {
	var empty_token model.Token
	token := mocks.GenerateToken()

	testCases := []struct {
		name           string
		expected       model.Token
		expectedSqlErr error
		tokenStr       string
	}{
		{
			name:           "return token",
			expected:       token,
			expectedSqlErr: nil,
			tokenStr:       "token.token.token",
		},
		{
			name:           "return error if not found",
			expected:       empty_token,
			expectedSqlErr: sql.ErrNoRows,
			tokenStr:       "error.error.error",
		},
	}

	for _, test := range testCases {
		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

		t.Run(test.name, func(t *testing.T) {
			columns := []string{
				"ID",
				"Uuid",
				"Token",
				"UserEmail",
				"Created",
			}
			expectedRow := sqlmock.
				NewRows(columns).
				AddRow(
					test.expected.ID,
					test.expected.Uuid,
					test.expected.Token,
					test.expected.UserEmail,
					test.expected.Created,
				)

			mock.
				ExpectQuery("SELECT ID, Uuid, Token, UserEmail, Created FROM token WHERE Token = ?").
				WithArgs(test.tokenStr).
				WillReturnRows(expectedRow).
				WillReturnError(test.expectedSqlErr)
			defer db.Close()

			repo := repository.NewAuthRepository(db)
			token, getErr := repo.GetToken(test.tokenStr)
			sqlErr := mock.ExpectationsWereMet()

			assert.AssertError(t, getErr, test.expectedSqlErr)
			assert.AssertError(t, sqlErr, nil)
			assert.AssertStruct[model.Token](t, token, test.expected)
		})
	}
}
