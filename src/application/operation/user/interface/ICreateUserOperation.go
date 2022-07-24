package operation_interface_user

import "github.com/andreis3/users-ms/src/application/operation/user/dto"

type ICreateUserOperation interface {
	Execute(userInput *dto.UserInputDTO) (*dto.UserOutPutDTO, error)
}
