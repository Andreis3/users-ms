package operation_interface_user

import "github.com/andreis3/users-ms/src/application/operation/user/dto"

//go:generate mockgen -package=mocks -destination ../../../../../tests/unit/mocks/GetAllUserOperationMock.go github.com/andreis3/users-ms/src/application/operation/user/interface IGetAllUserOperation
type IGetAllUserOperation interface {
	Execute() ([]*dto.UserOutPutDTO, error)
}
