package place_user

import (
	"github.com/google/uuid"

	"github.com/andreis3/users-ms/src/domain/entity"
)

type PlaceUserInput struct {
	Username  string `json:"user_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	CPF       string `json:"cpf" binding:"required"`
}

func (p *PlaceUserInput) ParserUserEntity() (*entity.User, error) {
	id := uuid.New().String()
	user := entity.NewUser(id, p.Username, p.Password, p.FirstName, p.LastName, p.Email, p.CPF)
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}
