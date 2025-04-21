package repositories

import (
	"context"
	"fmt"
	entities "task2/entities/user"
	"task2/exceptions"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (userRepository *UserRepository) Save(ctx context.Context, email, password string) (entities.User, exceptions.HttpError) {
	// Check email existed
	var emailExisted bool
	err := userRepository.db.GetContext(ctx, &emailExisted, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email)

	if err != nil {
		return entities.User{}, exceptions.NewInternal("")
	}

	if emailExisted {
		return entities.User{}, exceptions.NewConflict("Email existed")
	}

	query := "INSERT INTO users (email, password) VALUES($1, $2) RETURNING id"
	
	var id int64
	err = userRepository.db.QueryRowContext(ctx, query, email, password).Scan(&id)

	if err != nil {
		return entities.User{}, exceptions.NewInternal("")
	}

	var user entities.User
	err = userRepository.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return entities.User{}, exceptions.NewNotFound(fmt.Sprintf("Not found user with id = %d", id))
	}

	return user, exceptions.HttpError{}
}

func (userRepository *UserRepository) FindOne(context context.Context, id int64) (entities.User, exceptions.HttpError) {
	var user entities.User

	err := userRepository.db.GetContext(context, &user, "SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return entities.User{}, exceptions.NewNotFound(fmt.Sprintf("Not found user with id = %d", id))
	}

	return user, exceptions.HttpError{}
}

func (userRepository *UserRepository) FindByEmail(context context.Context, email string) (entities.User, exceptions.HttpError) {
	var user entities.User

	err := userRepository.db.GetContext(context, &user, "SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		return entities.User{}, exceptions.NewNotFound(fmt.Sprintf("Not found user with email = %s", email))
	}

	return user, exceptions.HttpError{}
}