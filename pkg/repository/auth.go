package repository

import (
	"database/sql"
	"tracker/pkg/errs"
	"tracker/pkg/model"
)

type AuthRepository interface {
	GetToken(token string) (model.Token, error)
	CreateToken(token model.Token) (error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) GetToken(tokenStr string) (model.Token, error) {
	query := `SELECT id, uuid, token, useremail, created FROM tokens WHERE token = $1`
	var token model.Token

	row := r.db.QueryRow(query, tokenStr)
	err := row.Scan(
		&token.ID,
		&token.Uuid,
		&token.Token,
		&token.UserEmail,
		&token.Created,
	)

	if err == sql.ErrNoRows {
		return token, errs.AuthTokenNotFound
	}

	if err != nil {
		return token, errs.Generic400Err
	}

	return token, nil
}

func (r *authRepository) CreateToken(token model.Token) error {
	query := `INSERT INTO tokens (token, useremail) VALUES ($1, $2)`
	_, err := r.db.Exec(query, token.Token, token.UserEmail)

	if err != nil {
		return err
	}

	return nil
}
