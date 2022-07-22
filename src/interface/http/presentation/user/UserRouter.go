package user_interface

import (
	"net/http"

	"github.com/andreis3/users-ms/src/infra/factory"
)

var userController IUserController
var databaseRepositoryFactory factory.DatabaseRepositoryFactory

func init() {
	userController = NewUserController(&databaseRepositoryFactory)
}

type UserRouter struct{}

func (u *UserRouter) UserRouter() []map[string]any {
	return []map[string]any{
		{
			"method": http.MethodPost,
			"path":   "/users",
			"handle": userController.Create,
		},
		{
			"method": http.MethodGet,
			"path":   "/users",
			"handle": userController.GetAll,
		},
		{
			"method": http.MethodGet,
			"path":   "/users/:id",
			"handle": userController.GetID,
		},
		{
			"method": http.MethodPut,
			"path":   "/users/:id",
			"handle": userController.Update,
		},
		{
			"method": http.MethodDelete,
			"path":   "/users/:id",
			"handle": userController.Delete,
		},
	}
}
