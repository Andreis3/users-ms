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

type UserRouter struct{}

func (u *UserRouter) UserRouter() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"method": http.MethodPost,
			"path":   "/users",
			"handle": userController.Create,
		},
		{
			"method": http.MethodGet,
			"path":   "/users",
			"handle": func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Hello World",
				})
			},
		},
	}
}
