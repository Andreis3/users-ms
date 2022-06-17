package http

import (
	"github.com/gin-gonic/gin"
)

type Http interface {
	On(method string, url string, handler func(c *gin.Context))
	Filter(filter func(c *gin.Context) bool)
	Listen(port string)
}
