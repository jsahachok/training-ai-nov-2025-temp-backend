package repositories

import "workshop4-backend/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id uint) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Update(id uint, user *entities.User) error
	Delete(id uint) error
	GetByEmail(email string) (*entities.User, error)
}
