package presentation

import userInterface "github.com/andreis3/users-ms/src/interface/http/presentation/user"

type IRegister interface {
	RegisterUser() userInterface.UserRouter
}
