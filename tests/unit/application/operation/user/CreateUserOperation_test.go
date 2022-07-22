package operation

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	operation "github.com/andreis3/users-ms/src/application/operation/user"
	"github.com/andreis3/users-ms/tests/unit/mocks"
)

var crtl *gomock.Controller

func Test_UserService(t *testing.T) {
	crtl = gomock.NewController(t)
	RegisterFailHandler(Fail)
	RunSpecs(t, "APPLICATION :: OPERATION :: USER :: CreateUserOperation Suite")
}

var _ = Describe("APPLICATION :: OPERATION :: USER :: CreateUserOperation", func() {
	var createUserOperation *operation.CreateUserOperation
	mockRepositoryFactory := mocks.NewMockIRepositoryFactory(crtl)
	mockUserRepository := mocks.NewMockIUserRepository(crtl)

})
