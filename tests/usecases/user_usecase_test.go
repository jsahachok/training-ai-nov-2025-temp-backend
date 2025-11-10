package usecases_test

import (
	"errors"
	"testing"
	"time"
	"workshop4-backend/domain/entities"
	"workshop4-backend/tests/mocks"
	"workshop4-backend/usecases"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	mockRepo    *mocks.MockUserRepository
	userUseCase *usecases.UserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.mockRepo = &mocks.MockUserRepository{}
	suite.userUseCase = usecases.NewUserUseCase(suite.mockRepo)
}

func (suite *UserUseCaseTestSuite) TestCreateUser_Success() {
	// Arrange
	req := &entities.UserCreateRequest{
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

	suite.mockRepo.On("GetByEmail", req.Email).Return(nil, nil)
	suite.mockRepo.On("Create", mock.MatchedBy(func(user *entities.User) bool {
		return user.Email == req.Email && user.FirstName == req.FirstName
	})).Return(nil)

	// Act
	user, err := suite.userUseCase.CreateUser(req)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), user)
	assert.Equal(suite.T(), req.FirstName, user.FirstName)
	assert.Equal(suite.T(), req.Email, user.Email)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestCreateUser_EmailAlreadyExists() {
	// Arrange
	req := &entities.UserCreateRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	existingUser := &entities.User{
		ID:    1,
		Email: req.Email,
	}

	suite.mockRepo.On("GetByEmail", req.Email).Return(existingUser, nil)

	// Act
	user, err := suite.userUseCase.CreateUser(req)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Contains(suite.T(), err.Error(), "user with this email already exists")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestGetUserByID_Success() {
	// Arrange
	userID := uint(1)
	expectedUser := &entities.User{
		ID:        userID,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	suite.mockRepo.On("GetByID", userID).Return(expectedUser, nil)

	// Act
	user, err := suite.userUseCase.GetUserByID(userID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), user)
	assert.Equal(suite.T(), expectedUser.ID, user.ID)
	assert.Equal(suite.T(), expectedUser.FirstName, user.FirstName)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestGetUserByID_InvalidID() {
	// Act
	user, err := suite.userUseCase.GetUserByID(0)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Contains(suite.T(), err.Error(), "invalid user ID")
}

func (suite *UserUseCaseTestSuite) TestGetUserByID_NotFound() {
	// Arrange
	userID := uint(999)
	suite.mockRepo.On("GetByID", userID).Return(nil, errors.New("user not found"))

	// Act
	user, err := suite.userUseCase.GetUserByID(userID)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestGetAllUsers_Success() {
	// Arrange
	expectedUsers := []*entities.User{
		{ID: 1, FirstName: "John", LastName: "Doe"},
		{ID: 2, FirstName: "Jane", LastName: "Smith"},
	}

	suite.mockRepo.On("GetAll").Return(expectedUsers, nil)

	// Act
	users, err := suite.userUseCase.GetAllUsers()

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), len(expectedUsers), len(users))
	assert.Equal(suite.T(), expectedUsers[0].FirstName, users[0].FirstName)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestUpdateUser_Success() {
	// Arrange
	userID := uint(1)
	newFirstName := "Jane"
	newEmail := "jane@example.com"

	existingUser := &entities.User{
		ID:        userID,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	req := &entities.UserUpdateRequest{
		FirstName: &newFirstName,
		Email:     &newEmail,
	}

	suite.mockRepo.On("GetByID", userID).Return(existingUser, nil)
	suite.mockRepo.On("GetByEmail", newEmail).Return(nil, nil)
	suite.mockRepo.On("Update", userID, mock.MatchedBy(func(user *entities.User) bool {
		return user.FirstName == newFirstName && user.Email == newEmail
	})).Return(nil)

	// Act
	user, err := suite.userUseCase.UpdateUser(userID, req)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), user)
	assert.Equal(suite.T(), newFirstName, user.FirstName)
	assert.Equal(suite.T(), newEmail, user.Email)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestDeleteUser_Success() {
	// Arrange
	userID := uint(1)
	existingUser := &entities.User{ID: userID, FirstName: "John"}

	suite.mockRepo.On("GetByID", userID).Return(existingUser, nil)
	suite.mockRepo.On("Delete", userID).Return(nil)

	// Act
	err := suite.userUseCase.DeleteUser(userID)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
