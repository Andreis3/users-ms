package repository

import (
	"github.com/andreis3/users-ms/src/domain/entity"
)

//go:generate mockgen -package=mocks -destination ../../../tests/unit/mocks/UserRepositoryMock.go github.com/andreis3/users-ms/src/domain/repository UserRepository

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
}
