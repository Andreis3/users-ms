package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	place_user "github.com/andreis3/users-ms/src/application/place-user"
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/infra/middleware/auth"
)

type RouterConfig struct {
	http              Http
	repositoryFactory factory.RepositoryFactory
	middleware        auth.MiddlewareInterface
}

func NewRouterConfig(http Http, repositoryFactory factory.RepositoryFactory, middleware auth.MiddlewareInterface) *RouterConfig {
	return &RouterConfig{
		http:              http,
		repositoryFactory: repositoryFactory,
		middleware:        middleware,
	}
}

func (routerConfig *RouterConfig) Build() {
	httpConfig := routerConfig.http
	repositoryFactory := routerConfig.repositoryFactory
	middleware := routerConfig.middleware

	httpConfig.Filter(func(c any) bool {
		ctx := c.(*gin.Context)
		ctx.Next()
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " - " +
			ctx.Request.Method + " " +
			ctx.Request.URL.Path + " - " +
			fmt.Sprintf("%v", ctx.Writer.Status()))
		return true
	})

	httpConfig.On(http.MethodPost, "/user", func(c any) {
		ctx := c.(*gin.Context)
		middleware.Middleware()(ctx)
		if ctx.Writer.Status() == http.StatusUnauthorized {
			return
		}
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
