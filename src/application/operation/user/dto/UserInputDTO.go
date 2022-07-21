package dto

import (
	"github.com/google/uuid"

	"github.com/andreis3/users-ms/src/domain/entity"
)

type UserInputDTO struct {
	Username  string `form:"user_name" json:"user_name" binding:"required" type:"string"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	CPF       string `form:"cpf" json:"cpf" binding:"required"`
}

func (p *UserInputDTO) ParserUserEntity() (*entity.User, error) {
	id := uuid.New().String()
	user := entity.NewUser(id, p.Username, p.Password, p.FirstName, p.LastName, p.Email, p.CPF)
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}
