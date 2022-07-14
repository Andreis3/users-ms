package operation

import (
	"time"

	"github.com/andreis3/users-ms/src/domain/entity"
)

type UserOutPutDTO struct {
	ID         string    `json:"id"`
	Username   string    `json:"user_name"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	CPF        string    `json:"cpf"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func ParserUserEntityOutput(userEntity *entity.User) *UserOutPutDTO {
	user := &UserOutPutDTO{
		ID:         userEntity.ID,
		Username:   userEntity.Username,
		FirstName:  userEntity.FirstName,
		LastName:   userEntity.LastName,
		Email:      userEntity.Email,
		CPF:        userEntity.CPF,
		CreatedAt:  userEntity.CreatedAt,
		ModifiedAt: userEntity.ModifiedAt,
	}
	return user
}
