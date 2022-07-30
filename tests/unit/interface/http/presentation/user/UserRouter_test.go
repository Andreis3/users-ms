package user_interfaces

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	user_interface "github.com/andreis3/users-ms/src/interface/http/presentation/user"
	"github.com/andreis3/users-ms/tests/unit/mocks"
)

var _ = Describe("UserRouter", func() {
	var userRouter *user_interface.UserRouter
	mockUserController := mocks.NewMockIUserController(ctrl)

	BeforeEach(func() {
		userRouter = user_interface.NewUserRouter(mockUserController)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should return the correct routes", func() {
		routerExpected := 5
		router := userRouter.UserRouter()
		Expect(len(router)).To(Equal(routerExpected))
	})
})
