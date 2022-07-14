package user_interface

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/andreis3/users-ms/src/infra/factory"
)

var userController *UserController
var memoryRepositoryFactory factory.MemoryRepositoryFactory

func init() {
	userController = NewUserController(memoryRepositoryFactory)
}

type UserRouter struct {
	Method string
	Path   string
	Handle func(ctx *gin.Context)
}

func (u *UserRouter) Register() []UserRouter {
	userRouters := []UserRouter{
		{
			Method: http.MethodPost,
			Path:   "/users",
			Handle: userController.Create,
		},
		{
			Method: http.MethodGet,
			Path:   "/users",
			Handle: func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Hello World",
				})
			},
		},
	}
	return userRouters
}
