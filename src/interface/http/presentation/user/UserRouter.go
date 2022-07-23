package user_interface

import (
	"net/http"
)

type UserRouter struct {
	userController IUserController
}

func NewUserRouter(userController IUserController) *UserRouter {
	return &UserRouter{
		userController: userController,
	}
}

func (r *UserRouter) UserRouter() []map[string]any {
	return []map[string]any{
		{
			"method": http.MethodPost,
			"path":   "/users",
			"handle": r.userController.Create,
		},
		{
			"method": http.MethodGet,
			"path":   "/users",
			"handle": r.userController.GetAll,
		},
		{
			"method": http.MethodGet,
			"path":   "/users/:id",
			"handle": r.userController.GetID,
		},
		{
			"method": http.MethodPut,
			"path":   "/users/:id",
			"handle": r.userController.Update,
		},
		{
			"method": http.MethodDelete,
			"path":   "/users/:id",
			"handle": r.userController.Delete,
		},
	}
}
