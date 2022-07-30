package operation_interface_user

import "github.com/andreis3/users-ms/src/application/operation/user/dto"

//go:generate mockgen -package=mocks -destination ../../../../../tests/unit/mocks/GetIdUserOperationMock.go github.com/andreis3/users-ms/src/application/operation/user/interface IGetIdUserOperation
type IGetIdUserOperation interface {
	Execute(id string) (*dto.UserOutPutDTO, error)
}
