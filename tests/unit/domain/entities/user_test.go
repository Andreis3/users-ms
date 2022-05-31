package entities

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/users-ms/domain/entities"
)

func TestUserEntity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test User Entity")
}

var _ = Describe("Validate User", Label("Entity"), func() {
	user, err := entities.NewUser(1,
		"test_username",
		"test_passworD1",
		"test_first_name",
		"test_last_name",
		"test_email@email.com",
		"06238411007")
	BeforeEach(func() {
		Expect(err).To(BeNil())
	})
	When("The user fields is valid", func() {
		It("Should return no error", func() {
			err := user.Validate()
			Expect(err).To(BeNil())
		})
	})
})
