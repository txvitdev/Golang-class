package repositories

import (
	"context"
	"fmt"
	entities "task2/internal/entities/user"
	"task2/internal/exceptions"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, exceptions.HttpError)
	CreateUser(ctx context.Context, user *entities.User) exceptions.HttpError
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) CreateUser(ctx context.Context, user *entities.User) exceptions.HttpError {
	query := `
		INSERT INTO users (id, email, password, full_name, is_email_verified)
		VALUES (:id, :email, :password, :full_name, :is_email_verified)
	`
	_, err := r.db.NamedExecContext(ctx, query, user)

	if err != nil {
		fmt.Println(err)
		return exceptions.NewInternal("")
	}
	return exceptions.HttpError{}
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*entities.User, exceptions.HttpError) {
	var user entities.User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE email=$1", email)

	if err != nil {
		return &user, exceptions.NewNotFound(fmt.Sprintf("Not found user with email = %s", email))
	}
	return &user, exceptions.HttpError{}
}
