package operation_interface_user

import "github.com/andreis3/users-ms/src/application/operation/user/dto"

//go:generate mockgen -package=mocks -destination ../../../../../tests/unit/mocks/UpdateUserOperationMock.go github.com/andreis3/users-ms/src/application/operation/user/interface IUpdateUserOperation
type IUpdateUserOperation interface {
	Execute(id string, userInput *dto.UserInputDTO) (*dto.UserOutPutDTO, error)
}
