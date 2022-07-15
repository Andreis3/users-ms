package _interface

import (
	"github.com/gin-gonic/gin"

	"github.com/andreis3/users-ms/src/interface/http/presentation"
	"github.com/andreis3/users-ms/src/interface/http/presentation/middleware/auth"
	userInterface "github.com/andreis3/users-ms/src/interface/http/presentation/user"
)

type Router struct {
	router         *gin.Engine
	registerRouter presentation.IRegisterRouter
	userRouter     userInterface.IUserRouter
}

func NewRouter(server *gin.Engine,
	registerRouter presentation.IRegisterRouter,
	userRouter userInterface.IUserRouter) *Router {
	return &Router{
		router:         server,
		registerRouter: registerRouter,
		userRouter:     userRouter,
	}
}

func (r *Router) ApiRouter() {
	auth := auth.NewAuthMiddleware("12345")
	userGroup := r.router.Group("/api/v1/")
	userGroup.Use(auth.Middleware())
	{
		r.registerRouter.Register(userGroup, r.userRouter.UserRouter())
	}
}
