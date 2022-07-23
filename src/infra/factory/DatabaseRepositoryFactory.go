package factory

import (
	"github.com/andreis3/users-ms/src/domain/repository"
	cockroach "github.com/andreis3/users-ms/src/infra/database/cockroachdb"
	"github.com/andreis3/users-ms/src/infra/repository/database"
)

type DatabaseRepositoryFactory struct{}

func NewDatabaseRepositoryFactory() *DatabaseRepositoryFactory {
	return &DatabaseRepositoryFactory{}
}

func (f *DatabaseRepositoryFactory) CreateUserRepository() repository.IUserRepository {
	return database.NewUserRepositoryDatabase(cockroach.GetInstance())
}
