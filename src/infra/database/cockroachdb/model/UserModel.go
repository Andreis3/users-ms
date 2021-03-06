package model

import "time"

type UserModel struct {
	ID         string    `json:"id" db:"id"`
	Username   string    `json:"user_name" db:"user_name"`
	Password   string    `json:"password" db:"password"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	Email      string    `json:"email" db:"email"`
	CPF        string    `json:"cpf" db:"cpf"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	ModifiedAt time.Time `json:"modified_at" db:"modified_at"`
}
