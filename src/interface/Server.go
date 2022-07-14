package _interface

import (
	"github.com/gin-gonic/gin"

	userInterface "github.com/andreis3/users-ms/src/interface/http/presentation/user"
)

type Server struct {
	registerRouter userInterface.UserRouter
	Server         *gin.Engine
}

func NewServer(registerUser userInterface.UserRouter) *Server {
	return &Server{
		Server:         gin.Default(),
		registerRouter: registerUser,
	}
}

func (s Server) Start() {
	s.Server = gin.Default()
	r := s.registerRouter.Register()
	for _, route := range r {
		s.Server.Handle(route.Method, route.Path, route.Handle)
	}

	s.Server.Run(":8080")
}
