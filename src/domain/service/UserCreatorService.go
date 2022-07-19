package service

import (
	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/repository"
)

type UserCreator struct {
	userRepository repository.IUserRepository
}

func NewUserCreator(repositoryFactory factory.IRepositoryFactory) *UserCreator {
	return &UserCreator{
		userRepository: repositoryFactory.CreateUserRepository(),
	}
}

func (u *UserCreator) Create(user *entity.User) (*entity.User, error) {
	userRepository, err := u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return userRepository, nil
}
