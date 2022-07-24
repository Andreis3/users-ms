package operation_user

import (
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/service"
)

type CreateUserOperation struct {
	userService service.IUserService
}

func NewCreateUserOperation(userService service.IUserService) *CreateUserOperation {
	return &CreateUserOperation{
		userService: userService,
	}
}

func (p *CreateUserOperation) Execute(userInput *dto.UserInputDTO) (*dto.UserOutPutDTO, error) {
	userEntity, err := userInput.ParserUserEntity()
	if err != nil {
		return nil, err
	}
	user, err := p.userService.CreateUser(userEntity)
	if err != nil {
		return nil, err
	}
	userOutput := dto.ParserUserEntityOutput(user)
	return userOutput, nil
}
