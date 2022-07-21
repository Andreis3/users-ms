package database

import (
	"fmt"
	"time"

	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/src/infra/database"
	"github.com/andreis3/users-ms/src/infra/database/model"
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
	sess := *u.db.Session()
	collection := sess.Collection("users")
	userResult := &model.UserModel{
		ID:         user.ID,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		CPF:        user.CPF,
		CreatedAt:  user.CreatedAt.String(),
		ModifiedAt: user.ModifiedAt.String(),
	}
	_, err := collection.Insert(userResult)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepositoryDatabase) Update(user *entity.User) (*entity.User, error) {
	userResult := &entity.User{
		ID:         user.ID,
		Username:   "test",
		Password:   "123456Ada$",
		FirstName:  "test",
		LastName:   "test",
		Email:      "test@mail.com",
		CPF:        "12345678912",
		CreatedAt:  time.Time{},
		ModifiedAt: time.Time{},
	}
	return userResult, nil
}

func (u UserRepositoryDatabase) FindByID(id string) (*entity.User, error) {
	userResult := &entity.User{
		ID:         id,
		Username:   "test",
		Password:   "123456Ada$",
		FirstName:  "test",
		LastName:   "test",
		Email:      "test@mail.com",
		CPF:        "12345678912",
		CreatedAt:  time.Time{},
		ModifiedAt: time.Time{},
	}
	return userResult, nil
}

func (u UserRepositoryDatabase) FindALL() ([]*entity.User, error) {
	usersResult := []*entity.User{
		{
			ID:         "1",
			Username:   "test",
			Password:   "123456Ada$",
			FirstName:  "test",
			LastName:   "test",
			Email:      "test@email.com",
			CPF:        "12345678912",
			CreatedAt:  time.Time{},
			ModifiedAt: time.Time{},
		},
		{
			ID:         "2",
			Username:   "test",
			Password:   "123456Ada$",
			FirstName:  "test",
			LastName:   "test",
			Email:      "test@email.com",
			CPF:        "12345678912",
			CreatedAt:  time.Time{},
			ModifiedAt: time.Time{},
		},
	}
	return usersResult, nil
}

func (u UserRepositoryDatabase) Delete(id string) error {
	fmt.Println("Deleting user-database with id: ", id)
	return nil
}
