package repository

import (
	"database/sql"
	"tracker/pkg/errs"
	"tracker/pkg/model"
)

type UserRepository interface {
	GetUser(id int) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	query := `SELECT id, uuid, email, password, username, created FROM users WHERE email = $1`
	var user model.User
	row := r.db.QueryRow(query, email)
	err := row.Scan(
		&user.ID,
		&user.Uuid,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.Created,
	)

	if err == sql.ErrNoRows {
		return user, errs.User404Err
	}

	if err != nil {
		return user, errs.Generic400Err
	}

	return user, nil
}

func (r *userRepository) GetUser(id int) (model.User, error) {
	query := `SELECT id, uuid, email, password, username, created FROM users WHERE id = $1`
	var user model.User
	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&user.ID,
		&user.Uuid,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.Created,
	)

	if err == sql.ErrNoRows {
		return user, errs.User404Err
	}

	if err != nil {
		return user, errs.Generic400Err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user model.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
