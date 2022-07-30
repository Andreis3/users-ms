//go:build unit
// +build unit

package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/src/domain/service"
	"github.com/andreis3/users-ms/tests/unit/mocks"
)

var crtl *gomock.Controller

func Test_UserService(t *testing.T) {
	crtl = gomock.NewController(t)
	RegisterFailHandler(Fail)
	RunSpecs(t, "DOMAIN :: SERVICE :: USER_SERVICE")
}

var _ = Describe("DOMAIN :: SERVICE :: CreateUser", func() {
	user := &entity.User{
		ID:        "any_id",
		Username:  "test_username",
		Password:  "test_password",
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Email:     "test_email@.com",
	}

	When("All fields are valid", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		BeforeEach(func() {
			mockUserRepository.EXPECT().Save(gomock.Any()).Return(user, nil)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})

		It("Should not return error", func() {
			userResult, err := userCreator.CreateUser(user)
			Expect(err).To(BeNil())
			Expect(userResult).To(Equal(user))
		})
	})

	When("Return a error", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		BeforeEach(func() {
			err := errors.New("username is required")
			mockUserRepository.EXPECT().Save(gomock.Any()).Return(nil, err)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})
		It("Should return error when username is empty", func() {
			userResult, err := userCreator.CreateUser(user)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("username is required"))
			Expect(userResult).To(BeNil())
		})
	})
})

var _ = Describe("DOMAIN :: SERVICE :: GetUserByID", func() {
	user := &entity.User{
		ID:        "any_id",
		Username:  "test_username",
		Password:  "test_password",
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Email:     "test_email@.com",
	}

	When("User exists", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		BeforeEach(func() {
			mockUserRepository.EXPECT().FindByID(gomock.Any()).Return(user, nil)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})

		AfterEach(func() {
			defer crtl.Finish()
		})

		It("Should not return error", func() {
			userResult, err := userCreator.GetUserByID("any_id")
			Expect(err).To(BeNil())
			Expect(userResult).To(Equal(user))
		})
	})

	When("User does not exist", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		BeforeEach(func() {
			err := errors.New("user not found")
			mockUserRepository.EXPECT().FindByID(gomock.Any()).Return(nil, err)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})
		It("Should return error when user does not exist", func() {
			userResult, err := userCreator.GetUserByID("any_id")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("user not found"))
			Expect(userResult).To(BeNil())
		})
	})

	When("ID is empty", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		BeforeEach(func() {
			err := errors.New("id is empty")
			mockUserRepository.EXPECT().FindByID(gomock.Any()).Return(nil, err)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})

		It("Should return error when id is empty", func() {
			userResult, err := userCreator.GetUserByID("")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("id is empty"))
			Expect(userResult).To(BeNil())
		})
	})
})

var _ = Describe("DOMAIN :: SERVICE :: GetAllUsers", func() {
	users := []*entity.User{
		{
			ID:        "any_id",
			Username:  "test_username",
			Password:  "test_password",
			FirstName: "test_first_name",
			LastName:  "test_last_name",
			Email:     "test_email@.com",
		},
		{
			ID:        "any_id_2",
			Username:  "test_username_2",
			Password:  "test_password_2",
			FirstName: "test_first_name_2",
			LastName:  "test_last_name_2",
			Email:     "test_email@.com_2",
		},
	}

	When("Users exists", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		BeforeEach(func() {
			mockUserRepository.EXPECT().FindALL().Return(users, nil)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})

		It("Should not return error", func() {
			userResult, err := userCreator.GetAllUsers()
			Expect(err).To(BeNil())
			Expect(userResult).To(Equal(users))
		})
	})

	When("Users does not exist", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		var user []*entity.User
		BeforeEach(func() {
			mockUserRepository.EXPECT().FindALL().Return(user, nil)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})

		It("Should return error when users does not exist", func() {
			userResult, err := userCreator.GetAllUsers()
			Expect(err).To(BeNil())
			Expect(userResult).To(BeNil())
		})
	})
})

var _ = Describe("DOMAIN :: SERVICE :: UpdateUser", func() {
	user := &entity.User{
		ID:        "any_id",
		Username:  "test_username",
		Password:  "test_password",
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Email:     "test_email@.com",
	}

	When("User exists", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		BeforeEach(func() {
			mockUserRepository.EXPECT().Update(gomock.Any(), user).Return(user, nil)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})

		It("Should not return error", func() {
			userResult, err := userCreator.UpdateUser("any_id", user)
			Expect(err).To(BeNil())
			Expect(userResult).To(Equal(user))
		})
	})
	When("User does not exist", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		BeforeEach(func() {
			err := errors.New("user not found")
			mockUserRepository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, err)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})
		It("Should return error when user does not exist", func() {
			userResult, err := userCreator.UpdateUser("any_id", user)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("user not found"))
			Expect(userResult).To(BeNil())
		})
	})
})

var _ = Describe("DOMAIN :: SERVICE :: DeleteUser", func() {
	When("User exists", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		BeforeEach(func() {
			mockUserRepository.EXPECT().Delete(gomock.Any()).Return(nil)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})
		It("Should not return error", func() {
			err := userCreator.DeleteUser("any_id")
			Expect(err).To(BeNil())
		})
	})
	When("User does not exist", func() {
		var userCreator *service.UserService
		mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
		mockUserRepository := mocks.NewMockIUserRepository(crtl)
		BeforeEach(func() {
			err := errors.New("user not found")
			mockUserRepository.EXPECT().Delete(gomock.Any()).Return(err)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserService(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})
		It("Should return error when user does not exist", func() {
			err := userCreator.DeleteUser("any_id")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("user not found"))
		})
	})
})
