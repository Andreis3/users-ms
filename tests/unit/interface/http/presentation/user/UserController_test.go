//go:build unit
// +build unit

package user_interfaces

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/users-ms/src/application/operation/user/dto"
	user_interface "github.com/andreis3/users-ms/src/interface/http/presentation/user"
	"github.com/andreis3/users-ms/tests/unit/mocks"
)

var _ = Describe("Create", func() {
	gin.SetMode(gin.TestMode)
	httpRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpRecorder)

	var userController *user_interface.UserController
	mockCreateOperation := mocks.NewMockICreateUserOperation(ctrl)
	mockGetAllOperation := mocks.NewMockIGetAllUserOperation(ctrl)
	mockGetIdOperation := mocks.NewMockIGetIdUserOperation(ctrl)
	mockUpdateOperation := mocks.NewMockIUpdateUserOperation(ctrl)
	mockDeleteOperation := mocks.NewMockIDeleteUserOperation(ctrl)

	userDto := &dto.UserOutPutDTO{
		ID:         "any_id",
		Username:   "test_username",
		FirstName:  "test_first_name",
		LastName:   "test_last_name",
		Email:      "test_email@.com",
		CPF:        "12345678912",
		CreatedAt:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		ModifiedAt: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
	}

	BeforeEach(func() {
		mockCreateOperation.EXPECT().Execute(gomock.Any()).Return(userDto, nil)

		userController = user_interface.NewUserController(
			mockCreateOperation,
			mockGetAllOperation,
			mockGetIdOperation,
			mockUpdateOperation,
			mockDeleteOperation,
		)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("Should create user", func() {
		var userInput dto.UserInputDTO

		jsonBytes := []byte(`{"user_name":"test_username",
		"password":"test_password",
		"first_name":"test_first_name",
		"last_name":"test_last_name",
		"cpf":"12345678912",
		"email":"test_email@.com"}`)

		json.Unmarshal(jsonBytes, &userInput)

		ginContext.Request = httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonBytes))
		userController.Create(ginContext)

		Expect(ginContext.Writer.Status()).To(Equal(http.StatusCreated))

		Expect(userDto).To(HaveField("ID", "any_id"))
		Expect(userDto).To(HaveField("CreatedAt", time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)))
		Expect(userDto).To(HaveField("ModifiedAt", time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC)))
		Expect(userInput.Username).To(Equal(userDto.Username))
		Expect(userInput.FirstName).To(Equal(userDto.FirstName))
		Expect(userInput.LastName).To(Equal(userDto.LastName))
		Expect(userInput.CPF).To(Equal(userDto.CPF))
		Expect(userInput.Email).To(Equal(userDto.Email))

	})
})

var _ = Describe("GetAll", func() {
	gin.SetMode(gin.TestMode)
	httpRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpRecorder)

	var userController *user_interface.UserController
	mockCreateOperation := mocks.NewMockICreateUserOperation(ctrl)
	mockGetAllOperation := mocks.NewMockIGetAllUserOperation(ctrl)
	mockGetIdOperation := mocks.NewMockIGetIdUserOperation(ctrl)
	mockUpdateOperation := mocks.NewMockIUpdateUserOperation(ctrl)
	mockDeleteOperation := mocks.NewMockIDeleteUserOperation(ctrl)

	userDto := &dto.UserOutPutDTO{
		ID:         "any_id",
		Username:   "test_username",
		FirstName:  "test_first_name",
		LastName:   "test_last_name",
		Email:      "test_email@.com",
		CPF:        "12345678912",
		CreatedAt:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		ModifiedAt: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
	}

	BeforeEach(func() {
		mockGetAllOperation.EXPECT().Execute().Return([]*dto.UserOutPutDTO{userDto}, nil)

		userController = user_interface.NewUserController(
			mockCreateOperation,
			mockGetAllOperation,
			mockGetIdOperation,
			mockUpdateOperation,
			mockDeleteOperation,
		)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("Should get all users", func() {
		userResult := map[string][]*dto.UserOutPutDTO{}
		ginContext.Request = httptest.NewRequest("GET", "/users", nil)
		userController.GetAll(ginContext)

		json.Unmarshal(httpRecorder.Body.Bytes(), &userResult)

		Expect(ginContext.Writer.Status()).To(Equal(http.StatusOK))

		Expect(userResult["users"][0].ID).To(Equal(userDto.ID))
		Expect(userResult["users"][0].CreatedAt).To(Equal(userDto.CreatedAt))
		Expect(userResult["users"][0].ModifiedAt).To(Equal(userDto.ModifiedAt))
		Expect(userResult["users"][0].Username).To(Equal(userDto.Username))
		Expect(userResult["users"][0].FirstName).To(Equal(userDto.FirstName))
		Expect(userResult["users"][0].LastName).To(Equal(userDto.LastName))
		Expect(userResult["users"][0].CPF).To(Equal(userDto.CPF))
		Expect(userResult["users"][0].Email).To(Equal(userDto.Email))
	})
})

var _ = Describe("GetById", func() {
	gin.SetMode(gin.TestMode)
	httpRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpRecorder)

	var userController *user_interface.UserController
	mockCreateOperation := mocks.NewMockICreateUserOperation(ctrl)
	mockGetAllOperation := mocks.NewMockIGetAllUserOperation(ctrl)
	mockGetIdOperation := mocks.NewMockIGetIdUserOperation(ctrl)
	mockUpdateOperation := mocks.NewMockIUpdateUserOperation(ctrl)
	mockDeleteOperation := mocks.NewMockIDeleteUserOperation(ctrl)

	userDto := &dto.UserOutPutDTO{
		ID:         "any_id",
		Username:   "test_username",
		FirstName:  "test_first_name",
		LastName:   "test_last_name",
		Email:      "test_email@.com",
		CPF:        "12345678912",
		CreatedAt:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		ModifiedAt: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
	}

	BeforeEach(func() {
		mockGetIdOperation.EXPECT().Execute(gomock.Any()).Return(userDto, nil)

		userController = user_interface.NewUserController(
			mockCreateOperation,
			mockGetAllOperation,
			mockGetIdOperation,
			mockUpdateOperation,
			mockDeleteOperation,
		)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("Should get user by id", func() {
		userResult := dto.UserOutPutDTO{}
		ginContext.Params = gin.Params{
			{Key: "id", Value: "any_id"},
		}
		ginContext.Request = httptest.NewRequest("GET", "/users/any_id", nil)
		userController.GetID(ginContext)
		json.Unmarshal(httpRecorder.Body.Bytes(), &userResult)

		Expect(ginContext.Writer.Status()).To(Equal(http.StatusOK))

		Expect(userResult.ID).To(Equal(userDto.ID))
		Expect(userResult.CreatedAt).To(Equal(userDto.CreatedAt))
		Expect(userResult.ModifiedAt).To(Equal(userDto.ModifiedAt))
		Expect(userResult.Username).To(Equal(userDto.Username))
		Expect(userResult.FirstName).To(Equal(userDto.FirstName))
		Expect(userResult.LastName).To(Equal(userDto.LastName))
		Expect(userResult.CPF).To(Equal(userDto.CPF))
		Expect(userResult.Email).To(Equal(userDto.Email))
	})
})

var _ = Describe("Update", func() {
	gin.SetMode(gin.TestMode)
	httpRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpRecorder)

	var userController *user_interface.UserController
	mockCreateOperation := mocks.NewMockICreateUserOperation(ctrl)
	mockGetAllOperation := mocks.NewMockIGetAllUserOperation(ctrl)
	mockGetIdOperation := mocks.NewMockIGetIdUserOperation(ctrl)
	mockUpdateOperation := mocks.NewMockIUpdateUserOperation(ctrl)
	mockDeleteOperation := mocks.NewMockIDeleteUserOperation(ctrl)

	userDto := &dto.UserOutPutDTO{
		ID:         "any_id",
		Username:   "update_username",
		FirstName:  "update_first_name",
		LastName:   "update_last_name",
		Email:      "test_email@.com",
		CPF:        "12345678912",
		CreatedAt:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		ModifiedAt: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
	}

	BeforeEach(func() {
		mockUpdateOperation.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(userDto, nil)

		userController = user_interface.NewUserController(
			mockCreateOperation,
			mockGetAllOperation,
			mockGetIdOperation,
			mockUpdateOperation,
			mockDeleteOperation,
		)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("Should update user", func() {
		userResult := dto.UserOutPutDTO{}
		ginContext.Params = gin.Params{
			{Key: "id", Value: "any_id"},
		}

		jsonBody := []byte(`{
			"user_name":"update_username",
			"password":"update_password",
			"first_name":"update_first_name",
			"last_name":"update_last_name",
			"email":"test_email@.com",
			"cpf":"12345678912"
		}`)

		ginContext.Request = httptest.NewRequest("PUT", "/users/any_id", bytes.NewBuffer(jsonBody))
		userController.Update(ginContext)
		json.Unmarshal(httpRecorder.Body.Bytes(), &userResult)

		Expect(ginContext.Writer.Status()).To(Equal(http.StatusOK))

		Expect(userResult.ID).To(Equal(userDto.ID))
		Expect(userResult.CreatedAt).To(Equal(userDto.CreatedAt))
		Expect(userResult.ModifiedAt).To(Equal(userDto.ModifiedAt))
		Expect(userResult.Username).To(Equal(userDto.Username))
		Expect(userResult.FirstName).To(Equal(userDto.FirstName))
		Expect(userResult.LastName).To(Equal(userDto.LastName))
		Expect(userResult.CPF).To(Equal(userDto.CPF))
		Expect(userResult.Email).To(Equal(userDto.Email))
	})
})

var _ = Describe("Delete", func() {
	gin.SetMode(gin.TestMode)
	httpRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpRecorder)

	var userController *user_interface.UserController
	mockCreateOperation := mocks.NewMockICreateUserOperation(ctrl)
	mockGetAllOperation := mocks.NewMockIGetAllUserOperation(ctrl)
	mockGetIdOperation := mocks.NewMockIGetIdUserOperation(ctrl)
	mockUpdateOperation := mocks.NewMockIUpdateUserOperation(ctrl)
	mockDeleteOperation := mocks.NewMockIDeleteUserOperation(ctrl)

	BeforeEach(func() {
		mockDeleteOperation.EXPECT().Execute(gomock.Any()).Return(nil)
		userController = user_interface.NewUserController(
			mockCreateOperation,
			mockGetAllOperation,
			mockGetIdOperation,
			mockUpdateOperation,
			mockDeleteOperation,
		)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("Should delete user", func() {
		ginContext.Params = gin.Params{
			{Key: "id", Value: "any_id"},
		}
		ginContext.Request = httptest.NewRequest(http.MethodDelete, "/users/any_id", nil)
		userController.Delete(ginContext)
		Expect(ginContext.Writer.Status()).To(Equal(http.StatusNoContent))
	})
})
