package _interface

import (
	"github.com/gin-gonic/gin"

	"github.com/andreis3/users-ms/src/interface/http/presentation"
	userInterface "github.com/andreis3/users-ms/src/interface/http/presentation/user"
)

type Server struct {
	Server     *gin.Engine
	router     presentation.IRegisterRouter
	userRouter userInterface.IUserRouter
}

func NewServer(userRouter userInterface.IUserRouter) *Server {
	return &Server{
		Server:     gin.Default(),
		userRouter: userRouter,
	}
}

func (s Server) Start() {
	s.router = presentation.NewRouterRegister(s.Server)
	s.router.Register(s.userRouter.UserRouter())
	s.Server.Run(":8080").Error()
}
