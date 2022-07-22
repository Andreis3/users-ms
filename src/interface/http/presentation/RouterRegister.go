package presentation

import (
	"github.com/gin-gonic/gin"
)

type RouterRegister struct{}

func NewRouterRegister() *RouterRegister {
	return &RouterRegister{}
}

func (r RouterRegister) Register(app *gin.RouterGroup, routers []map[string]interface{}) gin.HandlerFunc {
	for _, router := range routers {
		app.Handle(router["method"].(string), router["path"].(string), router["handle"].(func(ctx *gin.Context)))
	}
	return nil
}
