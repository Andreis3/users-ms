package user_interface

import "github.com/gin-gonic/gin"

//go:generate mockgen -package=mocks -destination ../../../../../tests/unit/mocks/UserControllerMock.go github.com/andreis3/users-ms/src/interface/http/presentation/user IUserController
type IUserController interface {
	Create(ctx *gin.Context)
	GetID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
