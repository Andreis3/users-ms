//go:build unit
// +build unit

package operation_user

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	operation "github.com/andreis3/users-ms/src/application/operation/user"
	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/tests/unit/mocks"
)

var _ = Describe("APPLICATION :: OPERATION :: USER :: GetAllUserOperation", Ordered, func() {
	entityUser := &entity.User{
		ID:        "any_id",
		Username:  "test_username",
		Password:  "test_password12F",
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Email:     "test_email@.com",
		CPF:       "12345678912",
	}

	When("All fields are valid", func() {
		var getAllUserOperation *operation.GetAllUserOperation
		var mockUserService *mocks.MockIUserService

		BeforeEach(func() {
			mockUserService = mocks.NewMockIUserService(ctrl)
			mockUserService.EXPECT().GetAllUsers().Return([]*entity.User{entityUser}, nil)
			getAllUserOperation = operation.NewGetAllUserOperation(mockUserService)
		})

		AfterEach(func() {
			defer ctrl.Finish()
		})

		It("Should not return error", func() {
			resultUser, err := getAllUserOperation.Execute()
			Expect(err).To(BeNil())
			Expect(resultUser).To(HaveLen(1))
		})
	})
})
