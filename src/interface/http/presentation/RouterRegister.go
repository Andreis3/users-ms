package presentation

import userInterface "github.com/andreis3/users-ms/src/interface/http/presentation/user"

type RouterRegister struct {
	userController userInterface.UserRouter
}

func (r RouterRegister) RegisterUser() []userInterface.UserRouter {
	return r.userController.Register()
}
