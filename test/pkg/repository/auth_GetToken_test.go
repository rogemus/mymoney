package repository_test

import (
	"database/sql"
	"testing"
	"tracker/pkg/model"
	"tracker/pkg/repository"
	assert "tracker/pkg/utils"
	errors "tracker/pkg/utils"
	mocks "tracker/test/pkg/mocks"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_AuthRepo_GetToken(t *testing.T) {
	var empty_token model.Token
	token := mocks.GenerateToken()

	testCases := []struct {
		name           string
		expected       model.Token
		expectedErr    error
		expectedSqlErr error
		tokenStr       string
	}{
		{
			name:           "return token",
			expected:       token,
			expectedErr:    nil,
			expectedSqlErr: nil,
			tokenStr:       "token.token.token",
		},
		{
			name:           "return error if not found",
			expected:       empty_token,
			expectedErr:    errors.AuthTokenNotFound,
			expectedSqlErr: sql.ErrNoRows,
			tokenStr:       "error.error.error",
		},
	}

	for _, test := range testCases {
		query := `SELECT ID, Uuid, Token, UserEmail, Created FROM token WHERE Token = "?"`
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
				ExpectQuery(query).
				WithArgs(test.tokenStr).
				WillReturnRows(expectedRow).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewAuthRepository(db)

			token, errGet := repo.GetToken(test.tokenStr)
			sqlErr := mock.ExpectationsWereMet()

			assert.AssertStruct[model.Token](t, token, test.expected)
			assert.AssertError(t, errGet, test.expectedErr)
			assert.AssertError(t, sqlErr, nil)
		})
	}
}
