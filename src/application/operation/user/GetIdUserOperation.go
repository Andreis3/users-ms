package operation

import (
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/service"
)

type GetIdUserOperation struct {
	repositoryFactory factory.IRepositoryFactory
}

func NewGetIdUserOperation(repositoryFactory factory.IRepositoryFactory) *GetIdUserOperation {
	return &GetIdUserOperation{
		repositoryFactory: repositoryFactory,
	}
}

func (p *GetIdUserOperation) Execute(id string) (*dto.UserOutPutDTO, error) {
	userService := service.NewUserService(p.repositoryFactory)
	user, err := userService.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	userOutput := dto.ParserUserEntityOutput(user)
	return userOutput, nil
}
