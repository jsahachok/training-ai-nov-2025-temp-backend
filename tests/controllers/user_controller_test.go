package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"workshop4-backend/domain/entities"
	"workshop4-backend/interface/controllers"
	"workshop4-backend/tests/mocks"
	"workshop4-backend/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	app            *fiber.App
	mockRepo       *mocks.MockUserRepository
	userController *controllers.UserController
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockRepo = &mocks.MockUserRepository{}
	userUseCase := usecases.NewUserUseCase(suite.mockRepo)
	suite.userController = controllers.NewUserController(userUseCase)

	suite.app = fiber.New()
	suite.app.Get("/users", suite.userController.GetAllUsers)
	suite.app.Get("/users/:id", suite.userController.GetUserByID)
	suite.app.Post("/users", suite.userController.CreateUser)
	suite.app.Put("/users/:id", suite.userController.UpdateUser)
	suite.app.Delete("/users/:id", suite.userController.DeleteUser)
}

func (suite *UserControllerTestSuite) TestGetAllUsers_Success() {
	// Arrange
	users := []*entities.User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com"},
		{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com"},
	}
	suite.mockRepo.On("GetAll").Return(users, nil)

	// Act
	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := suite.app.Test(req)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestGetUserByID_Success() {
	// Arrange
	user := &entities.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}
	suite.mockRepo.On("GetByID", uint(1)).Return(user, nil)

	// Act
	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, _ := suite.app.Test(req)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestGetUserByID_InvalidID() {
	// Act
	req := httptest.NewRequest("GET", "/users/invalid", nil)
	resp, _ := suite.app.Test(req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, resp.StatusCode)
}

func (suite *UserControllerTestSuite) TestCreateUser_Success() {
	// Arrange
	createReq := entities.UserCreateRequest{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		Phone:       "1234567890",
		DateOfBirth: "1990-01-01",
		Address:     "123 Main St",
		City:        "Bangkok",
		Country:     "Thailand",
		PostalCode:  "10110",
	}

	suite.mockRepo.On("GetByEmail", createReq.Email).Return(nil, nil)
	suite.mockRepo.On("Create", mock.MatchedBy(func(user *entities.User) bool {
		return user.Email == createReq.Email
	})).Return(nil)

	body, _ := json.Marshal(createReq)

	// Act
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := suite.app.Test(req)

	// Assert
	assert.Equal(suite.T(), http.StatusCreated, resp.StatusCode)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestCreateUser_ValidationError() {
	// Arrange
	createReq := entities.UserCreateRequest{
		FirstName: "John",
		// Missing required fields
	}

	body, _ := json.Marshal(createReq)

	// Act
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := suite.app.Test(req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, resp.StatusCode)
}

func (suite *UserControllerTestSuite) TestUpdateUser_Success() {
	// Arrange
	userID := uint(1)
	newFirstName := "Jane"
	updateReq := entities.UserUpdateRequest{
		FirstName: &newFirstName,
	}

	existingUser := &entities.User{
		ID:        userID,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	suite.mockRepo.On("GetByID", userID).Return(existingUser, nil)
	suite.mockRepo.On("Update", userID, mock.MatchedBy(func(user *entities.User) bool {
		return user.FirstName == newFirstName
	})).Return(nil)

	body, _ := json.Marshal(updateReq)

	// Act
	req := httptest.NewRequest("PUT", "/users/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := suite.app.Test(req)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestDeleteUser_Success() {
	// Arrange
	userID := uint(1)
	existingUser := &entities.User{ID: userID, FirstName: "John"}

	suite.mockRepo.On("GetByID", userID).Return(existingUser, nil)
	suite.mockRepo.On("Delete", userID).Return(nil)

	// Act
	req := httptest.NewRequest("DELETE", "/users/1", nil)
	resp, _ := suite.app.Test(req)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
