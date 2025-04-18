package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	entities "task2/entities/user"
	"task2/exceptions"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (userRepository *UserRepository) Save(ctx context.Context, email, password string) (entities.User, exceptions.HttpError) {
	query := "INSERT INTO users (email, password) VALUES($1, $2) RETURNING id"
	
	var id int64
	err := userRepository.db.QueryRowContext(ctx, query, email, password).Scan(&id)

	if err != nil {
		return entities.User{}, exceptions.HttpError{
			Code:    http.StatusInternalServerError,
			Message: "Error when create user",
		}
	}

	var user entities.User

	err = userRepository.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user)

	if err != nil {
		return entities.User{}, exceptions.HttpError{
			Code: http.StatusNotFound,
			Message: fmt.Sprintf("Not found user with id = %d", id),
		}
	}

	return user, exceptions.HttpError{}
}

func (userRepository *UserRepository) FindOne(id int64) (entities.User, exceptions.HttpError) {
	var user entities.User

	err := userRepository.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user)

	if err != nil {
		return entities.User{}, exceptions.HttpError{
			Code: http.StatusNotFound,
			Message: fmt.Sprintf("Not found user with id = %d", id),
		}
	}

	return user, exceptions.HttpError{}
}

// func (userRepository *UserRepository) FindAll() (entities.User, exceptions.HttpError) {
// 	query := "SELECT * FROM users ORDER BY id ASC"

// 	users := []*entities.User{}

// 	err := userRepository.db.QueryRowContext()
// }