//go:build unit
// +build unit

package operation_user_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	operationUser "github.com/andreis3/users-ms/src/application/operation/user"
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/tests/unit/mocks"
)

var _ = Describe("OPERATION :: USER :: CreateUserOperation", Ordered, func() {
	When("All fields are valid", func() {
		inputUserDTO := &dto.UserInputDTO{
			Username:  "test_username",
			Password:  "test_password12F",
			FirstName: "test_first_name",
			LastName:  "test_last_name",
			Email:     "test_email@.com",
			CPF:       "12345678912",
		}

		entityUser := &entity.User{
			ID:        "any_id",
			Username:  "test_username",
			Password:  "test_password12F",
			FirstName: "test_first_name",
			LastName:  "test_last_name",
			Email:     "test_email@.com",
			CPF:       "12345678912",
		}

		var createUserOperation *operationUser.CreateUserOperation
		mockUserService := mocks.NewMockIUserService(ctrl)
		BeforeEach(func() {
			mockUserService.EXPECT().CreateUser(gomock.Any()).Return(entityUser, nil)
			createUserOperation = operationUser.NewCreateUserOperation(mockUserService)
		})

		AfterEach(func() {
			defer ctrl.Finish()
		})

		It("Should not return error", func() {
			res, err := createUserOperation.Execute(inputUserDTO)
			Expect(err).To(BeNil())
			Expect(res.ID).To(BeEquivalentTo("any_id"))
			Expect(res.Username).To(Equal(entityUser.Username))
			Expect(res.FirstName).To(Equal(entityUser.FirstName))
			Expect(res.LastName).To(Equal(entityUser.LastName))
			Expect(res.Email).To(Equal(entityUser.Email))
			Expect(res.CPF).To(Equal(entityUser.CPF))

		})
	})

	When("Username is empty", func() {
		inputUserDTO := &dto.UserInputDTO{
			Username:  "",
			Password:  "test_password12F",
			FirstName: "test_first_name",
			LastName:  "test_last_name",
			Email:     "test_email@.com",
			CPF:       "12345678912",
		}

		entityUser := &entity.User{
			ID:        "any_id",
			Username:  "test_username",
			Password:  "test_password12F",
			FirstName: "test_first_name",
			LastName:  "test_last_name",
			Email:     "test_email@.com",
			CPF:       "12345678912",
		}

		mockUserService := mocks.NewMockIUserService(ctrl)
		var createUserOperation *operationUser.CreateUserOperation
		BeforeEach(func() {
			mockUserService.EXPECT().CreateUser(gomock.Any()).Return(entityUser, nil)
			createUserOperation = operationUser.NewCreateUserOperation(mockUserService)
		})

		AfterEach(func() {
			defer ctrl.Finish()
		})

		It("Should return error", func() {
			_, err := createUserOperation.Execute(inputUserDTO)
			Expect(err).To(HaveOccurred())
			Expect(err).ToNot(BeNil())
		})
	})

	When("Method CreateUser return error", func() {
		inputUserDTO := &dto.UserInputDTO{
			Username:  "test_username",
			Password:  "test_password12F",
			FirstName: "test_first_name",
			LastName:  "test_last_name",
			Email:     "test_email@.com",
			CPF:       "12345678912",
		}
		mockUserService := mocks.NewMockIUserService(ctrl)
		var createUserOperation *operationUser.CreateUserOperation

		BeforeEach(func() {
			errorTest := errors.New("any_error")
			mockUserService.EXPECT().CreateUser(gomock.Any()).Return(nil, errorTest)
			createUserOperation = operationUser.NewCreateUserOperation(mockUserService)
		})

		AfterEach(func() {
			defer ctrl.Finish()
		})

		It("Should return error", func() {
			_, err := createUserOperation.Execute(inputUserDTO)
			Expect(err.Error()).To(Equal("any_error"))
		})
	})
})
