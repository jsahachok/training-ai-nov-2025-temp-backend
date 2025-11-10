package repositories_test

import (
	"testing"
	"workshop4-backend/domain/entities"
	"workshop4-backend/domain/repositories"
	infrastructureModels "workshop4-backend/infrastructure/models"
	infrastructureRepo "workshop4-backend/infrastructure/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	db       *gorm.DB
	userRepo repositories.UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	// Create in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	// Auto migrate
	err = db.AutoMigrate(&infrastructureModels.UserModel{})
	suite.Require().NoError(err)

	suite.db = db
	suite.userRepo = infrastructureRepo.NewUserRepository(db)
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	// Clean up after each test
	suite.db.Exec("DELETE FROM users")
}

func (suite *UserRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	user := &entities.User{
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

	// Act
	err := suite.userRepo.Create(user)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), user.ID)
	assert.NotZero(suite.T(), user.CreatedAt)
	assert.NotZero(suite.T(), user.UpdatedAt)
}

func (suite *UserRepositoryTestSuite) TestGetByID_Success() {
	// Arrange
	user := &entities.User{
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
	err := suite.userRepo.Create(user)
	suite.Require().NoError(err)

	// Act
	foundUser, err := suite.userRepo.GetByID(user.ID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), foundUser)
	assert.Equal(suite.T(), user.ID, foundUser.ID)
	assert.Equal(suite.T(), user.FirstName, foundUser.FirstName)
	assert.Equal(suite.T(), user.Email, foundUser.Email)
}

func (suite *UserRepositoryTestSuite) TestGetByID_NotFound() {
	// Act
	user, err := suite.userRepo.GetByID(999)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Contains(suite.T(), err.Error(), "user not found")
}

func (suite *UserRepositoryTestSuite) TestGetAll_Success() {
	// Arrange
	user1 := &entities.User{
		FirstName: "John", LastName: "Doe", Email: "john@example.com",
		Phone: "1234567890", DateOfBirth: "1990-01-01", Address: "123 Main St",
		City: "Bangkok", Country: "Thailand", PostalCode: "10110",
	}
	user2 := &entities.User{
		FirstName: "Jane", LastName: "Smith", Email: "jane@example.com",
		Phone: "0987654321", DateOfBirth: "1991-02-02", Address: "456 Oak St",
		City: "Bangkok", Country: "Thailand", PostalCode: "10111",
	}

	err := suite.userRepo.Create(user1)
	suite.Require().NoError(err)
	err = suite.userRepo.Create(user2)
	suite.Require().NoError(err)

	// Act
	users, err := suite.userRepo.GetAll()

	// Assert
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), users, 2)
}

func (suite *UserRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	user := &entities.User{
		FirstName: "John", LastName: "Doe", Email: "john@example.com",
		Phone: "1234567890", DateOfBirth: "1990-01-01", Address: "123 Main St",
		City: "Bangkok", Country: "Thailand", PostalCode: "10110",
	}
	err := suite.userRepo.Create(user)
	suite.Require().NoError(err)

	// Update user
	user.FirstName = "Johnny"
	user.LastName = "Updated"

	// Act
	err = suite.userRepo.Update(user.ID, user)

	// Assert
	assert.NoError(suite.T(), err)

	// Verify update
	updatedUser, err := suite.userRepo.GetByID(user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Johnny", updatedUser.FirstName)
	assert.Equal(suite.T(), "Updated", updatedUser.LastName)
}

func (suite *UserRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	user := &entities.User{
		FirstName: "John", LastName: "Doe", Email: "john@example.com",
		Phone: "1234567890", DateOfBirth: "1990-01-01", Address: "123 Main St",
		City: "Bangkok", Country: "Thailand", PostalCode: "10110",
	}
	err := suite.userRepo.Create(user)
	suite.Require().NoError(err)

	// Act
	err = suite.userRepo.Delete(user.ID)

	// Assert
	assert.NoError(suite.T(), err)

	// Verify deletion (soft delete)
	_, err = suite.userRepo.GetByID(user.ID)
	assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGetByEmail_Success() {
	// Arrange
	user := &entities.User{
		FirstName: "John", LastName: "Doe", Email: "john@example.com",
		Phone: "1234567890", DateOfBirth: "1990-01-01", Address: "123 Main St",
		City: "Bangkok", Country: "Thailand", PostalCode: "10110",
	}
	err := suite.userRepo.Create(user)
	suite.Require().NoError(err)

	// Act
	foundUser, err := suite.userRepo.GetByEmail(user.Email)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), foundUser)
	assert.Equal(suite.T(), user.Email, foundUser.Email)
}

func (suite *UserRepositoryTestSuite) TestGetByEmail_NotFound() {
	// Act
	user, err := suite.userRepo.GetByEmail("nonexistent@example.com")

	// Assert
	assert.NoError(suite.T(), err) // No error for not found
	assert.Nil(suite.T(), user)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
