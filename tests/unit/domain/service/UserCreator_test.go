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

func TestUserCreator(t *testing.T) {
	crtl = gomock.NewController(t)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain :: Service :: UserCreator Suite")
}

var _ = Describe("DOMAIN :: SERVICE :: USERCREATOR", func() {
	var userCreator *service.UserCreator
	mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
	mockUserRepository := mocks.NewMockIUserRepository(crtl)
	user := &entity.User{
		ID:        "any_id",
		Username:  "test_username",
		Password:  "test_password",
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Email:     "test_email@.com",
	}

	When("All fields are valid", func() {
		BeforeEach(func() {
			mockUserRepository.EXPECT().Save(gomock.Any()).Return(user, nil)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserCreator(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})

		It("Should not return error", func() {
			userResult, err := userCreator.Create(user)
			Expect(err).To(BeNil())
			Expect(userResult).To(Equal(user))
		})
	})

	When("Return a error", func() {
		BeforeEach(func() {
			err := errors.New("username is required")
			mockUserRepository.EXPECT().Save(gomock.Any()).Return(nil, err)
			mockRepositoryFactory.EXPECT().CreateUserRepository().Return(mockUserRepository)
			userCreator = service.NewUserCreator(mockRepositoryFactory)
		})
		AfterEach(func() {
			defer crtl.Finish()
		})
		It("Should return error when username is empty", func() {
			userResult, err := userCreator.Create(user)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("username is required"))
			Expect(userResult).To(BeNil())
		})
	})
})
