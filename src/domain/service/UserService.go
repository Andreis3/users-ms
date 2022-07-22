package service

import (
	"errors"

	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/repository"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(repositoryFactory factory.IRepositoryFactory) *UserService {
	return &UserService{
		userRepository: repositoryFactory.CreateUserRepository(),
	}
}

func (u *UserService) CreateUser(user *entity.User) (*entity.User, error) {
	userRepository, err := u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return userRepository, nil
}

func (u *UserService) GetUserByID(id string) (*entity.User, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	userRepository, err := u.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return userRepository, nil
}

func (u *UserService) GetAllUsers() ([]*entity.User, error) {
	userRepository, err := u.userRepository.FindALL()
	if err != nil {
		return nil, err
	}
	return userRepository, nil
}

func (u *UserService) UpdateUser(id string, user *entity.User) (*entity.User, error) {
	userResult, err := u.userRepository.Update(id, user)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}

func (u *UserService) DeleteUser(id string) error {
	err := u.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
