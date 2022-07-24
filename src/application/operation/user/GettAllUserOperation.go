package operationUser

import (
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/service"
)

type GetAllUserOperation struct {
	userService service.IUserService
}

func NewGetAllUserOperation(userService service.IUserService) *GetAllUserOperation {
	return &GetAllUserOperation{
		userService: userService,
	}
}

func (p *GetAllUserOperation) Execute() ([]*dto.UserOutPutDTO, error) {
	users, err := p.userService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var userOutput []*dto.UserOutPutDTO
	for _, user := range users {
		userOutput = append(userOutput, dto.ParserUserEntityOutput(user))
	}
	return userOutput, nil
}
