package place_user

import (
	"github.com/andreis3/users-ms/src/domain/factory"
	"github.com/andreis3/users-ms/src/domain/service"
)

type PlaceUser struct {
	repositoryFactory factory.RepositoryFactory
}

func NewPlaceUser(repositoryFactory factory.RepositoryFactory) *PlaceUser {
	return &PlaceUser{
		repositoryFactory: repositoryFactory,
	}
}

func (p *PlaceUser) Execute(userInput *PlaceUserInput) (*PlaceUserOutPut, error) {
	userRepository := service.NewUserCreator(p.repositoryFactory)
	userEntity, err := userInput.ParserUserEntity()
	if err != nil {
		return nil, err
	}
	user, err := userRepository.Create(userEntity)
	if err != nil {
		return nil, err
	}
	userOutput := ParserUserEntityOutput(user)
	return userOutput, nil
}
