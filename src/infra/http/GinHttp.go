package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHttp struct {
	app *gin.Engine
}

func NewGinHttp() GinHttp {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,authentication")
		c.Next()
	})

	router.OPTIONS("/*path", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	return GinHttp{
		app: router,
	}
}

func (h GinHttp) Filter(filter func(c any) bool) {
	h.app.Use(func(c *gin.Context) {
		if filter(c) {
			c.Next()
			//fmt.Println("Filter:", c.Request.Method, c.Request.URL.Path, c.Writer.Status())

		} else {
			c.AbortWithStatus(c.Writer.Status())
			return
		}
	})
}

func (h GinHttp) On(method string, url string, fn func(c any)) {
	h.app.Handle(method, url, func(c *gin.Context) {
		fn(c)
	})
}

func (h GinHttp) Listen(port string) {
	fmt.Println("Start server in port:", port)
	h.app.Run(":" + port)
}
