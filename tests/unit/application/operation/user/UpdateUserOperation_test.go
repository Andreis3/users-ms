//go:build unit
// +build unit

package operation_user

import (
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	operationUser "github.com/andreis3/users-ms/src/application/operation/user"
	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/tests/unit/mocks"
)

var _ = Describe("UpdateUserOperation", func() {
	When("All fields are valid", func() {
		userInputDTO := &dto.UserInputDTO{
			Username:  "test_username",
			Password:  "test_password12F",
			FirstName: "test_first_name",
			LastName:  "test_last_name",
			Email:     "test_email@.com",
			CPF:       "12345678912",
		}

		entityUser := &entity.User{
			ID:         "any_id",
			Username:   "test_username",
			Password:   "test_password12F",
			FirstName:  "test_first_name",
			LastName:   "test_last_name",
			Email:      "test_email@.com",
			CPF:        "12345678912",
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}

		var updateUserOperation *operationUser.UpdateUserOperation
		mockUserService := mocks.NewMockIUserService(ctrl)
		BeforeEach(func() {
			mockUserService.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(entityUser, nil)
			updateUserOperation = operationUser.NewUpdateUserOperation(mockUserService)
		})

		AfterEach(func() {
			defer ctrl.Finish()
		})

		It("Should not return error", func() {
			result, err := updateUserOperation.Execute("any_id", userInputDTO)
			Expect(err).To(BeNil())
			Expect(result.ID).To(BeEquivalentTo("any_id"))
		})
	})
})
