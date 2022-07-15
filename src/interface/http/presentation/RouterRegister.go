package presentation

import (
	"github.com/gin-gonic/gin"
)

type RouterRegister struct {
	Server *gin.Engine
}

func NewRouterRegister(server *gin.Engine) *RouterRegister {
	return &RouterRegister{
		Server: server,
	}
}

func (r RouterRegister) Register(router []map[string]interface{}) []map[string]interface{} {
	for _, route := range router {
		r.Server.Handle(route["method"].(string), route["path"].(string), route["handle"].(func(ctx *gin.Context)))
	}
	return router
}
