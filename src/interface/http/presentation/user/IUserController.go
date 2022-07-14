package user_interface

import "github.com/gin-gonic/gin"

type IUserController interface {
	Create(ctx *gin.Context)
}
