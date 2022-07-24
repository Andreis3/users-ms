package operation_interface_user

import "github.com/andreis3/users-ms/src/application/operation/user/dto"

type IGetIdUserOperation interface {
	Execute(id string) (*dto.UserOutPutDTO, error)
}
