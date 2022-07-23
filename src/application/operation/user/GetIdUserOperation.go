package operation

import (
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/service"
)

type GetIdUserOperation struct {
	userService service.IUserService
}

func NewGetIdUserOperation(userService service.IUserService) *GetIdUserOperation {
	return &GetIdUserOperation{
		userService: userService,
	}
}

func (p *GetIdUserOperation) Execute(id string) (*dto.UserOutPutDTO, error) {
	user, err := p.userService.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	userOutput := dto.ParserUserEntityOutput(user)
	return userOutput, nil
}
