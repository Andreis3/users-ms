package user_interface

import "github.com/gin-gonic/gin"

type IUserController interface {
	Create(ctx *gin.Context)
	GetID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
