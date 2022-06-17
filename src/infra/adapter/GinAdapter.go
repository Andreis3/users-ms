package adapter

import "github.com/gin-gonic/gin"

type GinAdapter struct {
	Handler func(c *gin.Context)
}
