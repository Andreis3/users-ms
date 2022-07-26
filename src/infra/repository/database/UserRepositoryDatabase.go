package database

import (
	"time"

	"github.com/google/uuid"

	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/src/infra/database"
	"github.com/andreis3/users-ms/src/infra/database/cockroachdb/model"
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
		ID:         uuid.New().String(),
		Username:   user.Username,
		FirstName:  user.FirstName,
		Password:   user.Password,
		LastName:   user.LastName,
		Email:      user.Email,
		CPF:        user.CPF,
		CreatedAt:  user.CreatedAt,
		ModifiedAt: user.ModifiedAt,
	}
	res, err := collection.Insert(userResult)
	if err != nil {
		return nil, err
	}
	user.ID = res.ID().(string)
	return user, nil
}

func (u UserRepositoryDatabase) Update(id string, user *entity.User) (*entity.User, error) {
	sess := *u.db.Session()
	collection := sess.Collection("users")
	userModel := &model.UserModel{
		ID:         id,
		Username:   user.Username,
		FirstName:  user.FirstName,
		Password:   user.Password,
		LastName:   user.LastName,
		Email:      user.Email,
		CPF:        user.CPF,
		CreatedAt:  user.CreatedAt,
		ModifiedAt: time.Now(),
	}
	err := collection.Find("id", id).Update(userModel)
	if err != nil {
		return nil, err
	}
	user.ID = userModel.ID
	user.Username = userModel.Username
	user.FirstName = userModel.FirstName
	user.LastName = userModel.LastName
	user.Email = userModel.Email
	user.CPF = userModel.CPF
	user.CreatedAt = userModel.CreatedAt
	user.ModifiedAt = userModel.ModifiedAt

	return user, nil
}

func (u UserRepositoryDatabase) FindByID(id string) (*entity.User, error) {
	sess := *u.db.Session()
	collection := sess.Collection("users")
	var userModel *model.UserModel
	err := collection.Find("id", id).One(&userModel)
	if err != nil {
		return nil, err
	}
	userResult := &entity.User{
		ID:         userModel.ID,
		Username:   userModel.Username,
		FirstName:  userModel.FirstName,
		LastName:   userModel.LastName,
		Email:      userModel.Email,
		CPF:        userModel.CPF,
		CreatedAt:  userModel.CreatedAt,
		ModifiedAt: userModel.ModifiedAt,
	}
	return userResult, nil
}

func (u UserRepositoryDatabase) FindALL() ([]*entity.User, error) {
	sess := *u.db.Session()
	var userModels []*model.UserModel
	err := sess.Collection("users").Find().All(&userModels)
	if err != nil {
		return nil, err
	}
	var userResults []*entity.User
	for _, userModel := range userModels {
		userResult := &entity.User{
			ID:         userModel.ID,
			Username:   userModel.Username,
			FirstName:  userModel.FirstName,
			LastName:   userModel.LastName,
			Email:      userModel.Email,
			CPF:        userModel.CPF,
			CreatedAt:  userModel.CreatedAt,
			ModifiedAt: userModel.ModifiedAt,
		}
		userResults = append(userResults, userResult)
	}
	return userResults, nil
}

func (u UserRepositoryDatabase) Delete(id string) error {
	sess := *u.db.Session()
	collection := sess.Collection("users")
	err := collection.Find("id", id).Delete()
	if err != nil {
		return err
	}
	return nil
}
