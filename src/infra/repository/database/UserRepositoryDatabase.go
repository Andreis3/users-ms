package database

import (
	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/src/infra/database"
)

type UserRepositoryDatabase struct {
	db database.IDatabase
}

func NewUserRepositoryDatabase(db database.IDatabase) UserRepositoryDatabase {
	return UserRepositoryDatabase{
		db: db,
	}
}

func (u UserRepositoryDatabase) Save(user *entity.User) (*entity.User, error) {
	userResult := u.db.One(user)
	return userResult.(*entity.User), nil
}

func (u UserRepositoryDatabase) Update(user *entity.User) (*entity.User, error) {
	userResult := u.db.One(user)
	return userResult.(*entity.User), nil
}

func (u UserRepositoryDatabase) FindByID(id string) (*entity.User, error) {
	userResult := u.db.One(id)
	return userResult.(*entity.User), nil
}

func (u UserRepositoryDatabase) FindALL() ([]*entity.User, error) {
	usersResult := u.db.Many(nil)
	return usersResult.([]*entity.User), nil
}

func (u UserRepositoryDatabase) Delete(id string) error {
	u.db.None(id)
	return nil
}
