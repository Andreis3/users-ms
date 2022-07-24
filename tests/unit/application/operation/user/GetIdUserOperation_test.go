//go:build unit
// +build unit

package operation_user_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	operation "github.com/andreis3/users-ms/src/application/operation/user"
	"github.com/andreis3/users-ms/src/domain/entity"
	"github.com/andreis3/users-ms/tests/unit/mocks"
)

var _ = Describe("APPLICATION :: OPERATION :: USER :: GetIdUserOperation", Ordered, func() {
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
		var getIdUserOperation *operation.GetIdUserOperation
		var mockUserService *mocks.MockIUserService

		BeforeEach(func() {
			mockUserService = mocks.NewMockIUserService(ctrl)
			mockUserService.EXPECT().GetUserByID(gomock.Any()).Return(entityUser, nil)
			getIdUserOperation = operation.NewGetIdUserOperation(mockUserService)
		})

		AfterEach(func() {
			defer ctrl.Finish()
		})

		It("Should not return error", func() {
			_, err := getIdUserOperation.Execute("any_id")
			Expect(err).To(BeNil())
		})
	})
})
