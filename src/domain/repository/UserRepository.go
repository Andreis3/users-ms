package repository

import (
	"github.com/andreis3/users-ms/src/domain/entity"
)

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
}
