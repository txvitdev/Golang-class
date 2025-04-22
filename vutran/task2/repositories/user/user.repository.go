package repositories

import (
	"context"
	"fmt"
	dto "task2/dtos/auth"
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

func (userRepository *UserRepository) Save(ctx context.Context, signUpDto *dto.SignUpDto) (entities.User, exceptions.HttpError) {
	query := "INSERT INTO users (email, password) VALUES($1, $2) RETURNING id"
	
	var id int64
	err := userRepository.db.QueryRowContext(ctx, query, signUpDto.Email, signUpDto.Password).Scan(&id)

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

func (userRepository *UserRepository) IsEmailExisted(ctx context.Context, email string) (bool, exceptions.HttpError) {
	var isExisted bool

	err := userRepository.db.GetContext(ctx, &isExisted, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email)

	if err != nil {
		return false, exceptions.NewInternal("")
	}

	return isExisted, exceptions.HttpError{}
} 

