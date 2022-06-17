package factory

import (
	"github.com/andreis3/users-ms/src/domain/repository"
	"github.com/andreis3/users-ms/src/infra/repository/memory"
)

type MemoryRepositoryFactory struct {
	userRepository repository.UserRepository
}

func (f MemoryRepositoryFactory) CreateUserRepository() repository.UserRepository {
	f.userRepository = memory.GetInstance()
	return f.userRepository
}
