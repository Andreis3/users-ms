package service

import (
	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/repository"
)

type UserCreator struct {
	userRepository repository.UserRepository
}

func NewUserCreator(repositoryFactory factory.RepositoryFactory) *UserCreator {
	return &UserCreator{
		userRepository: repositoryFactory.CreateUserRepository(),
	}
}

func (u *UserCreator) Create(user *entity.User) (*entity.User, error) {
	userRepository, _ := u.userRepository.Save(user)
	return userRepository, nil
}
