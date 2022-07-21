package user_interface

import (
	"net/http"

	"github.com/gin-gonic/gin"

	operation "github.com/andreis3/users-ms/src/application/operation/user"
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/factory"
)

type UserController struct {
	repositoryFactory factory.IRepositoryFactory
}

func NewUserController(repositoryFactory factory.IRepositoryFactory) *UserController {
	return &UserController{
		repositoryFactory: repositoryFactory,
	}
}

func (u *UserController) Create(ctx *gin.Context) {
	var userInput dto.UserInputDTO
	ctx.ShouldBindJSON(&userInput)
	userOutput, err := operation.NewCreateUserOperation(u.repositoryFactory).Execute(&userInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, userOutput)
}

func (u *UserController) GetID(ctx *gin.Context) {
	id := ctx.Param("id")
	userOutput, err := operation.NewGetIdUserOperation(u.repositoryFactory).Execute(id)
	if err != nil && userOutput == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, userOutput)
}

func (u *UserController) GetAll(ctx *gin.Context) {
	usersOutput, err := operation.NewGetAllUserOperation(u.repositoryFactory).Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": usersOutput})
}

func (u *UserController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var userInput dto.UserInputDTO
	ctx.ShouldBindJSON(&userInput)
	userOutput, err := operation.NewUpdateUserOperation(u.repositoryFactory).Execute(id, &userInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, userOutput)
}

func (u *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := operation.NewDeleteUserOperation(u.repositoryFactory).Execute(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
