package memory

import (
	"sync"

	"github.com/andreis3/users-ms/src/domain/entity"
)

var (
	instance       *UserRepositoryMemory
	once           sync.Once
	memoryDatabase []any
)

type UserRepositoryMemory struct{}

func GetInstance() *UserRepositoryMemory {
	once.Do(func() {
		instance = &UserRepositoryMemory{}
	})
	return instance
}

func (u UserRepositoryMemory) Save(user *entity.User) (*entity.User, error) {
	memoryDatabase = append(memoryDatabase, user)
	return user, nil
}
