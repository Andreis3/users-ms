package main

import (
	"github.com/gin-gonic/gin"

	_interface "github.com/andreis3/users-ms/src/interface"
	"github.com/andreis3/users-ms/src/interface/http/presentation"
	userInterface "github.com/andreis3/users-ms/src/interface/http/presentation/user"
)

var router *_interface.Router
var userRouter *userInterface.UserRouter
var registerRouter *presentation.RouterRegister

func main() {
	app := gin.Default()
	registerRouter = presentation.NewRouterRegister()
	router = _interface.NewRouter(app, registerRouter, userRouter)
	server := _interface.NewServer(app, router)
	server.Start()
}
