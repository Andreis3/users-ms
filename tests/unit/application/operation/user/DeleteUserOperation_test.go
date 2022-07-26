//go:build unit
// +build unit

package operation_user

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	operationUser "github.com/andreis3/users-ms/src/application/operation/user"
	"github.com/andreis3/users-ms/tests/unit/mocks"
)

var _ = Describe("DeleteUserOperation", func() {
	When("All fields are valid", func() {
		var deleteUserOperation *operationUser.DeleteUserOperation
		mockUserService := mocks.NewMockIUserService(ctrl)
		BeforeEach(func() {
			mockUserService.EXPECT().DeleteUser(gomock.Any()).Return(nil)
			deleteUserOperation = operationUser.NewDeleteUserOperation(mockUserService)
		})

		AfterEach(func() {
			defer ctrl.Finish()
		})

		It("Should not return error", func() {
			err := deleteUserOperation.Execute("any_id")
			Expect(err).To(BeNil())
		})
	})
})
