package operation

import (
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/service"
)

type DeleteUserOperation struct {
	repositoryFactory factory.IRepositoryFactory
}

func NewDeleteUserOperation(repositoryFactory factory.IRepositoryFactory) *DeleteUserOperation {
	return &DeleteUserOperation{
		repositoryFactory: repositoryFactory,
	}
}

func (p *DeleteUserOperation) Execute(id string) error {
	userService := service.NewUserService(p.repositoryFactory)
	err := userService.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
