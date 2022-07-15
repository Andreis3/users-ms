package user_interface

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/andreis3/users-ms/src/infra/factory"
)

var userController IUserController
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
					"message": "Implement method GET",
				})
			},
		},
		{
			"method": http.MethodPut,
			"path":   "/users/:id",
			"handle": func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Implement method PUT",
				})
			},
		},
		{
			"method": http.MethodDelete,
			"path":   "/users/:id",
			"handle": func(ctx *gin.Context) {
				ctx.JSON(http.StatusNoContent, gin.H{
					"message": "Implement method DELETE",
				})
			},
		},
	}
}
