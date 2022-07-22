package factory

import "github.com/andreis3/users-ms/src/domain/repository"

//go:generate mockgen -package=mocks -destination ../../../tests/unit/mocks/RepositoryFactoryMock.go github.com/andreis3/users-ms/src/domain/factory IRepositoryFactory

type IRepositoryFactory interface {
	CreateUserRepository() repository.IUserRepository
}
