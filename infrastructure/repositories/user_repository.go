package infrastructure

import (
	"errors"
	"workshop4-backend/domain/entities"
	"workshop4-backend/domain/repositories"
	infrastructureModels "workshop4-backend/infrastructure/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *entities.User) error {
	userModel := infrastructureModels.FromEntity(user)

	result := r.db.Create(userModel)
	if result.Error != nil {
		return result.Error
	}

	// Update the entity with the generated ID and timestamps
	*user = *userModel.ToEntity()
	return nil
}

func (r *userRepository) GetByID(id uint) (*entities.User, error) {
	var userModel infrastructureModels.UserModel

	result := r.db.First(&userModel, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return userModel.ToEntity(), nil
}

func (r *userRepository) GetAll() ([]*entities.User, error) {
	var userModels []infrastructureModels.UserModel

	result := r.db.Find(&userModels)
	if result.Error != nil {
		return nil, result.Error
	}

	users := make([]*entities.User, len(userModels))
	for i, model := range userModels {
		users[i] = model.ToEntity()
	}

	return users, nil
}

func (r *userRepository) Update(id uint, user *entities.User) error {
	userModel := infrastructureModels.FromEntity(user)
	userModel.ID = id

	result := r.db.Save(userModel)
	if result.Error != nil {
		return result.Error
	}

	// Update the entity with the updated timestamps
	*user = *userModel.ToEntity()
	return nil
}

func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&infrastructureModels.UserModel{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepository) GetByEmail(email string) (*entities.User, error) {
	var userModel infrastructureModels.UserModel

	result := r.db.Where("email = ?", email).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil without error if not found
		}
		return nil, result.Error
	}

	return userModel.ToEntity(), nil
}
