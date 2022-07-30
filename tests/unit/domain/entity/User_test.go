//go:build unit
// +build unit

package entity

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/users-ms/src/domain/entity"
)

func Test_UserEntity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DOMAIN :: ENTITY :: USER")
}

var _ = Describe("DOMAIN :: ENTITY :: USER", func() {
	var user *entity.User
	var err error
	BeforeEach(func() {
		user = entity.NewUser(
			"test_username",
			"test_passworD1",
			"test_first_name",
			"test_last_name",
			"test_email@email.com",
			"06238411007")
		Expect(err).To(BeNil())
	})

	When("All fields are valid", func() {
		It("Should not return error", func() {
			err := user.Validate()
			Expect(err).To(BeNil())
		})
	})

	When("The name is empty", func() {
		It("Should return error when username is empty", func() {
			user.Username = ""
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("username is required"))
		})
	})

	When("The password is invalid type", func() {
		It("Should return error when password is empty", func() {
			user.Password = ""
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("password is required"))
		})

		It("Should return error when password contain less than 8 characters", func() {
			user.Password = "1234567"
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("password must be at least 8 characters"))
		})

		It("Should return error when password not contain characters special", func() {
			user.Password = "1234567890123456dD"
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("password must contain at least one special character"))
		})

		It("Should return error when password not contain characters lowercase", func() {
			user.Password = "1234567890123456DD"
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("password must contain at least one lower case letter"))
		})

		It("Should return error when password not contain characters uppercase", func() {
			user.Password = "1234567890123456dd"
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("password must contain at least one upper case letter"))
		})

		It("Should return error when password not contain characters number", func() {
			user.Password = "asdfghjkB@"
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("password must contain at least one number"))
		})
	})

	When("The CPF is invalid", func() {
		It("Should return error when CPF is empty", func() {
			user.CPF = ""
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("CPF is required"))
		})

		It("Should return error when CPF is invalid contains letter", func() {
			user.CPF = "1234567890A"
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("cpf is invalid"))
		})

		It("Should return error when CPF is invalid contains character special invalid", func() {
			user.CPF = "123.456.789$07"
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("cpf is invalid"))
		})
	})

	When("The CPF is valid", func() {
		It("Should not return error when CPF is valid", func() {
			user.CPF = "12345678909"
			err := user.Validate()
			Expect(err).To(BeNil())
		})

		It("Should not return error when CPF is valid with character special accepts", func() {
			user.CPF = "123.456.789-09"
			err := user.Validate()
			Expect(err).To(BeNil())
		})
	})

	When("The email is invalid", func() {
		It("Should return error when email is empty", func() {
			user.Email = ""
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("email is required"))
		})

		It("Should return error when email not contains .com", func() {
			user.Email = "teste@teste"
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("email must contain .com"))
		})

		It("Should return error when email not contains @", func() {
			user.Email = "teste.com"
			err := user.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("email must contain @"))
		})
	})
})
