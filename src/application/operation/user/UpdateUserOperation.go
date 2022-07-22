package operation

import (
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/service"
)

type UpdateUserOperation struct {
	repositoryFactory factory.IRepositoryFactory
}

func NewUpdateUserOperation(repositoryFactory factory.IRepositoryFactory) *UpdateUserOperation {
	return &UpdateUserOperation{
		repositoryFactory: repositoryFactory,
	}
}

func (p *UpdateUserOperation) Execute(id string, userInput *dto.UserInputDTO) (*entity.User, error) {
	userService := service.NewUserService(p.repositoryFactory)
	user, err := userInput.ParserUserEntity()
	if err != nil {
		return nil, err
	}
	userResult, err := userService.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}
