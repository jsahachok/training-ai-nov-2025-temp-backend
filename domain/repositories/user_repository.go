package repositories

import "workshop4-backend/domain/entities"

type TransferRepository interface {
	Create(transfer *entities.Transfer) error
	GetByID(id uint) (*entities.Transfer, error)
	GetByUserID(userID uint) ([]*entities.Transfer, error)
	GetTransferHistory(userID uint, limit, offset int) ([]*entities.Transfer, int, error)
	UpdateStatus(id uint, status string) error
	BeginTransaction() interface{}
	CommitTransaction(tx interface{}) error
	RollbackTransaction(tx interface{}) error
	CreateWithTransaction(tx interface{}, transfer *entities.Transfer) error
}

type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id uint) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Update(id uint, user *entities.User) error
	Delete(id uint) error
	GetByEmail(email string) (*entities.User, error)
	UpdatePoints(userID uint, points float64) error
	UpdatePointsWithTransaction(tx interface{}, userID uint, points float64) error
}
