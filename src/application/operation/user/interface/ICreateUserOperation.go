package operation_interface_user

import "github.com/andreis3/users-ms/src/application/operation/user/dto"

//go:generate mockgen -package=mocks -destination ../../../../../tests/unit/mocks/CreateUserOperationMock.go github.com/andreis3/users-ms/src/application/operation/user/interface ICreateUserOperation
type ICreateUserOperation interface {
	Execute(userInput *dto.UserInputDTO) (*dto.UserOutPutDTO, error)
}
