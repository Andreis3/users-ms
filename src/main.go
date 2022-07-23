package main

import (
	"github.com/gin-gonic/gin"

	operation "github.com/andreis3/users-ms/src/application/operation/user"
	"github.com/andreis3/users-ms/src/domain/service"
	"github.com/andreis3/users-ms/src/infra/factory"
	_interface "github.com/andreis3/users-ms/src/interface"
	"github.com/andreis3/users-ms/src/interface/http/presentation"
	userInterface "github.com/andreis3/users-ms/src/interface/http/presentation/user"
)

var (
	app            *gin.Engine
	router         *_interface.Router
	registerRouter *presentation.RouterRegister

	createUserOperation *operation.CreateUserOperation
	getIdUserOperation  *operation.GetIdUserOperation
	getAllUserOperation *operation.GetAllUserOperation
	updateUserOperation *operation.UpdateUserOperation
	deleteUserOperation *operation.DeleteUserOperation

	userController *userInterface.UserController

	userRouter *userInterface.UserRouter

	server *_interface.Server
)

func init() {
	app = gin.Default()

	databaseRepositoryFactory := factory.NewDatabaseRepositoryFactory()

	userService := service.NewUserService(databaseRepositoryFactory)

	createUserOperation = operation.NewCreateUserOperation(userService)
	getIdUserOperation = operation.NewGetIdUserOperation(userService)
	getAllUserOperation = operation.NewGetAllUserOperation(userService)
	updateUserOperation = operation.NewUpdateUserOperation(userService)
	deleteUserOperation = operation.NewDeleteUserOperation(userService)

	userController = userInterface.NewUserController(createUserOperation,
		getAllUserOperation, getIdUserOperation,
		updateUserOperation, deleteUserOperation)

	userRouter = userInterface.NewUserRouter(userController)
	registerRouter = presentation.NewRouterRegister()

	router = _interface.NewRouter(app, registerRouter, userRouter)
	server = _interface.NewServer(app, router)
}

func main() {
	server.Start()
}
