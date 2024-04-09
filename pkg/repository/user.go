package repository

import (
	"database/sql"
	"tracker/pkg/model"
	errors "tracker/pkg/utils"
)

type UserRepository interface {
	GetUser(id int) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (int64, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	query := "SELECT ID, Uuid, Email, Password, Username, Created FROM user WHERE Email = ?"
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
		return user, errors.User404Err
	}

	if err != nil {
		return user, errors.Generic400Err
	}

	return user, nil
}

func (r *userRepository) GetUser(id int) (model.User, error) {
	query := "SELECT ID, Uuid, Email, Created FROM user WHERE ID = ?"
	var user model.User
	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&user.ID,
		&user.Uuid,
		&user.Email,
		&user.Created,
	)

	if err == sql.ErrNoRows {
		return user, errors.User404Err
	}

	if err != nil {
		return user, errors.Generic400Err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user model.User) (int64, error) {
	query := "INSERT INTO user (Username, Email, Password) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, user.Username, user.Email, user.Password)

	if err != nil {
		return -1, errors.Generic400Err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, errors.Generic400Err
	}

	return id, nil
}
