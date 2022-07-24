package operation_interface_user

import "github.com/andreis3/users-ms/src/application/operation/user/dto"

type IUpdateUserOperation interface {
	Execute(id string, userInput *dto.UserInputDTO) (*dto.UserOutPutDTO, error)
}
