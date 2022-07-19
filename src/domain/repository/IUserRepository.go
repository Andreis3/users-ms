package repository

import (
	"github.com/andreis3/users-ms/src/domain/entity"
)

//go:generate mockgen -package=mocks -destination ../../../tests/unit/mocks/UserRepositoryMock.go github.com/andreis3/users-ms/src/domain/repository IUserRepository

type IUserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
	FindALL() ([]*entity.User, error)
	Delete(id string) error
}
