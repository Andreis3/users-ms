package user_interface

import (
	"net/http"

	"github.com/gin-gonic/gin"

	operation "github.com/andreis3/users-ms/src/application/operation/user"
	"github.com/andreis3/users-ms/src/domain/factory"
)

type UserController struct {
	repositoryFactory factory.RepositoryFactory
}

func NewUserController(repositoryFactory factory.RepositoryFactory) *UserController {
	return &UserController{
		repositoryFactory: repositoryFactory,
	}
}

func (u *UserController) Create(ctx *gin.Context) {
	var userInput operation.UserInputDTO
	ctx.ShouldBindJSON(&userInput)
	userOutput, err := operation.NewUserOperation(u.repositoryFactory).Execute(&userInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user": userOutput,
	})
}
