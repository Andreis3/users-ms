package operation

import (
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/service"
)

type CreateUserOperation struct {
	repositoryFactory factory.IRepositoryFactory
}

func NewUserOperation(repositoryFactory factory.IRepositoryFactory) *CreateUserOperation {
	return &CreateUserOperation{
		repositoryFactory: repositoryFactory,
	}
}

func (p *CreateUserOperation) Execute(userInput *dto.UserInputDTO) (*dto.UserOutPutDTO, error) {
	userRepository := service.NewUserCreator(p.repositoryFactory)
	userEntity, err := userInput.ParserUserEntity()
	if err != nil {
		return nil, err
	}
	user, err := userRepository.Create(userEntity)
	if err != nil {
		return nil, err
	}
	userOutput := dto.ParserUserEntityOutput(user)
	return userOutput, nil
}
