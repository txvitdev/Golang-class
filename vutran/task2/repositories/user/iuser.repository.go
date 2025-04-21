package repositories

import entities "task2/entities/user"

type IUserRepository interface {
	FindOne() entities.User
	FindAll() []entities.User
	Save() entities.User
}