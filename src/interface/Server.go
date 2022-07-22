package _interface

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Server *gin.Engine
	router IRouter
}

func NewServer(server *gin.Engine, router IRouter) *Server {
	return &Server{
		Server: server,
		router: router,
	}
}

func (s Server) Start() {
	s.router.ApiRouter()
	s.Server.Run(":3000").Error()
}
