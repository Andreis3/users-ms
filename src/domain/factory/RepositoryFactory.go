package factory

import "github.com/andreis3/users-ms/src/domain/repository"

type RepositoryFactory interface {
	CreateUserRepository() repository.UserRepository
}
