package user_interface

import (
	"net/http"

	"github.com/gin-gonic/gin"

	operation "github.com/andreis3/users-ms/src/application/operation/user"
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
)

type UserController struct {
	createUserOperation operation.ICreateUserOperation
	getAllUserOperation operation.IGetAllUserOperation
	getIdUserOperation  operation.IGetIdUserOperation
	updateUserOperation operation.IUpdateUserOperation
	deleteUserOperation operation.IDeleteUserOperation
}

func NewUserController(
	createUserOperation operation.ICreateUserOperation,
	getAllUserOperation operation.IGetAllUserOperation,
	getIdUserOperation operation.IGetIdUserOperation,
	updateUserOperation operation.IUpdateUserOperation,
	deleteUserOperation operation.IDeleteUserOperation,
) *UserController {
	return &UserController{
		createUserOperation: createUserOperation,
		getAllUserOperation: getAllUserOperation,
		getIdUserOperation:  getIdUserOperation,
		updateUserOperation: updateUserOperation,
		deleteUserOperation: deleteUserOperation,
	}
}

func (u *UserController) Create(ctx *gin.Context) {
	var userInput dto.UserInputDTO
	ctx.ShouldBindJSON(&userInput)
	userOutput, err := u.createUserOperation.Execute(&userInput)
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
	userOutput, err := u.getIdUserOperation.Execute(id)
	if err != nil && userOutput == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, userOutput)
}

func (u *UserController) GetAll(ctx *gin.Context) {
	usersOutput, err := u.getAllUserOperation.Execute()
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
	userOutput, err := u.updateUserOperation.Execute(id, &userInput)
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
	err := u.deleteUserOperation.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
