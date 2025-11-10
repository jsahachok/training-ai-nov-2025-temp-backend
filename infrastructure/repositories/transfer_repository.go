package infrastructure

import (
	"errors"
	"workshop4-backend/domain/entities"
	"workshop4-backend/domain/repositories"
	infrastructureModels "workshop4-backend/infrastructure/models"

	"gorm.io/gorm"
)

type transferRepository struct {
	db *gorm.DB
}

func NewTransferRepository(db *gorm.DB) repositories.TransferRepository {
	return &transferRepository{
		db: db,
	}
}

func (r *transferRepository) Create(transfer *entities.Transfer) error {
	transferModel := infrastructureModels.FromTransferEntity(transfer)

	result := r.db.Create(transferModel)
	if result.Error != nil {
		return result.Error
	}

	// Update the entity with the generated ID and timestamps
	*transfer = *transferModel.ToEntity()
	return nil
}

func (r *transferRepository) GetByID(id uint) (*entities.Transfer, error) {
	var transferModel infrastructureModels.TransferModel

	result := r.db.First(&transferModel, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("transfer not found")
		}
		return nil, result.Error
	}

	return transferModel.ToEntity(), nil
}

func (r *transferRepository) GetByUserID(userID uint) ([]*entities.Transfer, error) {
	var transferModels []infrastructureModels.TransferModel

	result := r.db.Where("from_user_id = ? OR to_user_id = ?", userID, userID).
		Order("created_at DESC").Find(&transferModels)
	if result.Error != nil {
		return nil, result.Error
	}

	transfers := make([]*entities.Transfer, len(transferModels))
	for i, model := range transferModels {
		transfers[i] = model.ToEntity()
	}

	return transfers, nil
}

func (r *transferRepository) GetTransferHistory(userID uint, limit, offset int) ([]*entities.Transfer, int, error) {
	var transferModels []infrastructureModels.TransferModel
	var total int64

	// Count total records
	countResult := r.db.Model(&infrastructureModels.TransferModel{}).
		Where("from_user_id = ? OR to_user_id = ?", userID, userID).Count(&total)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	// Get paginated records
	result := r.db.Where("from_user_id = ? OR to_user_id = ?", userID, userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transferModels)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	transfers := make([]*entities.Transfer, len(transferModels))
	for i, model := range transferModels {
		transfers[i] = model.ToEntity()
	}

	return transfers, int(total), nil
}

func (r *transferRepository) UpdateStatus(id uint, status string) error {
	result := r.db.Model(&infrastructureModels.TransferModel{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("transfer not found")
	}

	return nil
}

func (r *transferRepository) BeginTransaction() interface{} {
	return r.db.Begin()
}

func (r *transferRepository) CommitTransaction(tx interface{}) error {
	return tx.(*gorm.DB).Commit().Error
}

func (r *transferRepository) RollbackTransaction(tx interface{}) error {
	return tx.(*gorm.DB).Rollback().Error
}

func (r *transferRepository) CreateWithTransaction(tx interface{}, transfer *entities.Transfer) error {
	transferModel := infrastructureModels.FromTransferEntity(transfer)

	result := tx.(*gorm.DB).Create(transferModel)
	if result.Error != nil {
		return result.Error
	}

	// Update the entity with the generated ID and timestamps
	*transfer = *transferModel.ToEntity()
	return nil
}
