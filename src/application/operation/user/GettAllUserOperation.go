package operation

import (
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/service"
)

type GetAllUserOperation struct {
	repositoryFactory factory.IRepositoryFactory
}

func NewGetAllUserOperation(repositoryFactory factory.IRepositoryFactory) *GetAllUserOperation {
	return &GetAllUserOperation{
		repositoryFactory: repositoryFactory,
	}
}

func (p *GetAllUserOperation) Execute() ([]*dto.UserOutPutDTO, error) {
	userService := service.NewUserService(p.repositoryFactory)
	users, err := userService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var userOutput []*dto.UserOutPutDTO
	for _, user := range users {
		userOutput = append(userOutput, dto.ParserUserEntityOutput(user))
	}
	return userOutput, nil
}
