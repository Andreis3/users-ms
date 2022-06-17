package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	place_user "github.com/andreis3/users-ms/src/application/place-user"
	"github.com/andreis3/users-ms/src/domain/factory"
)

type RouterConfig struct {
	http              Http
	repositoryFactory factory.RepositoryFactory
}

func NewRouterConfig(http GinHttp, repositoryFactory factory.RepositoryFactory) *RouterConfig {
	return &RouterConfig{
		http:              http,
		repositoryFactory: repositoryFactory,
	}
}

func (routerConfig *RouterConfig) Build() {
	httpConfig := routerConfig.http
	repositoryFactory := routerConfig.repositoryFactory

	httpConfig.Filter(func(c *gin.Context) bool {
		c.Next()
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " - " +
			c.Request.Method + " " +
			c.Request.URL.Path + " - " +
			fmt.Sprintf("%v", c.Writer.Status()))
		return true
	})

	httpConfig.On("POST", "/user", func(ctx *gin.Context) {
		var userInput place_user.PlaceUserInput
		if err := ctx.ShouldBindJSON(&userInput); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userRepository := place_user.NewPlaceUser(repositoryFactory)
		user, err := userRepository.Execute(&userInput)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, user)
		return
	})
}
