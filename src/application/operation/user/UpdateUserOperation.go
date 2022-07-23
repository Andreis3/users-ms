package operation

import (
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/service"
)

type UpdateUserOperation struct {
	userService service.IUserService
}

func NewUpdateUserOperation(userService service.IUserService) *UpdateUserOperation {
	return &UpdateUserOperation{
		userService: userService,
	}
}

func (p *UpdateUserOperation) Execute(id string, userInput *dto.UserInputDTO) (*dto.UserOutPutDTO, error) {
	user, err := userInput.ParserUserEntity()
	if err != nil {
		return nil, err
	}
	userResult, err := p.userService.UpdateUser(id, user)
	userRes := dto.ParserUserEntityOutput(userResult)
	if err != nil {
		return nil, err
	}
	return userRes, nil
}
