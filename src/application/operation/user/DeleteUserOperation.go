package operation_user

import (
	"github.com/andreis3/users-ms/src/domain/service"
)

type DeleteUserOperation struct {
	userService service.IUserService
}

func NewDeleteUserOperation(userService service.IUserService) *DeleteUserOperation {
	return &DeleteUserOperation{
		userService: userService,
	}
}

func (p *DeleteUserOperation) Execute(id string) error {
	err := p.userService.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
