package factory

import (
	"github.com/andreis3/users-ms/src/domain/repository"
	database2 "github.com/andreis3/users-ms/src/infra/database"
	"github.com/andreis3/users-ms/src/infra/repository/database"
)

type DatabaseRepositoryFactory struct{}

func (f *DatabaseRepositoryFactory) CreateUserRepository() repository.IUserRepository {
	return database.NewUserRepositoryDatabase(database2.GetInstance())
}
