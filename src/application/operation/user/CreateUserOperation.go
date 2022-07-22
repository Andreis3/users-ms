package operation

import (
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/service"
)

type CreateUserOperation struct {
	repositoryFactory factory.IRepositoryFactory
}

func NewCreateUserOperation(repositoryFactory factory.IRepositoryFactory) *CreateUserOperation {
	return &CreateUserOperation{
		repositoryFactory: repositoryFactory,
	}
}

func (p *CreateUserOperation) Execute(userInput *dto.UserInputDTO) (*dto.UserOutPutDTO, error) {
	userService := service.NewUserService(p.repositoryFactory)
	userEntity, err := userInput.ParserUserEntity()
	if err != nil {
		return nil, err
	}
	user, err := userService.CreateUser(userEntity)
	if err != nil {
		return nil, err
	}
	userOutput := dto.ParserUserEntityOutput(user)
	return userOutput, nil
}
