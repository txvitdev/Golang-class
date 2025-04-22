package services

import (
	"context"
	entities "task2/entities/user"
	"task2/exceptions"
	repositories "task2/repositories/user"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (UserService *UserService) FindOne(ctx context.Context, id int64) (entities.User, exceptions.HttpError) {
	return UserService.userRepository.FindOne(ctx, id)
}