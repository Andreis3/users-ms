package operation_user_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/users-ms/tests/unit/mocks"
)

var ctrl *gomock.Controller
var mockUserService *mocks.MockIUserService

func TestUser(t *testing.T) {
	ctrl = gomock.NewController(t)
	RegisterFailHandler(Fail)
	RunSpecs(t, "OPERATION :: TESTS :: USER :: CreateUserOperation :: GetIdUserOperation")
}
