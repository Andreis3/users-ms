package presentation

import "github.com/gin-gonic/gin"

type IRegisterRouter interface {
	Register(app *gin.RouterGroup, router []map[string]interface{}) gin.HandlerFunc
}
