package service

import (
	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/repository"
)

type UserCreator struct {
	userRepository repository.IUserRepository
}

func NewUserService(repositoryFactory factory.IRepositoryFactory) *UserCreator {
	return &UserCreator{
		userRepository: repositoryFactory.CreateUserRepository(),
	}
}

func (u *UserCreator) CreateUser(user *entity.User) (*entity.User, error) {
	userRepository, err := u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return userRepository, nil
}

func (u *UserCreator) GetUserByID(id string) (*entity.User, error) {
	userRepository, err := u.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return userRepository, nil
}

func (u *UserCreator) GetAllUsers() ([]*entity.User, error) {
	userRepository, err := u.userRepository.FindALL()
	if err != nil {
		return nil, err
	}
	return userRepository, nil
}

func (u *UserCreator) UpdateUser(id string, user *entity.User) (*entity.User, error) {
	userResult, err := u.userRepository.Update(id, user)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}

func (u *UserCreator) DeleteUser(id string) error {
	err := u.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
