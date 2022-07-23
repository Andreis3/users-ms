package service

import "github.com/andreis3/users-ms/src/domain/entity"

type IUserService interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	GetAllUsers() ([]*entity.User, error)
	UpdateUser(id string, user *entity.User) (*entity.User, error)
	DeleteUser(id string) error
}
