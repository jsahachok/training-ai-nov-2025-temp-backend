package infrastructure

import (
	"time"
	"workshop4-backend/domain/entities"

	"gorm.io/gorm"
)

type TransferModel struct {
	ID          uint   `gorm:"primaryKey"`
	FromUserID  uint   `gorm:"not null"`
	ToUserID    uint   `gorm:"not null"`
	Points      int    `gorm:"not null"`
	Description string `gorm:"default:''"`
	Status      string `gorm:"not null;default:'pending'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (TransferModel) TableName() string {
	return "transfers"
}

func (t *TransferModel) ToEntity() *entities.Transfer {
	var deletedAt *time.Time
	if t.DeletedAt.Valid {
		deletedAt = &t.DeletedAt.Time
	}

	return &entities.Transfer{
		ID:          t.ID,
		FromUserID:  t.FromUserID,
		ToUserID:    t.ToUserID,
		Points:      t.Points,
		Description: t.Description,
		Status:      t.Status,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}

func FromTransferEntity(transfer *entities.Transfer) *TransferModel {
	model := &TransferModel{
		ID:          transfer.ID,
		FromUserID:  transfer.FromUserID,
		ToUserID:    transfer.ToUserID,
		Points:      transfer.Points,
		Description: transfer.Description,
		Status:      transfer.Status,
		CreatedAt:   transfer.CreatedAt,
		UpdatedAt:   transfer.UpdatedAt,
	}

	if transfer.DeletedAt != nil {
		model.DeletedAt = gorm.DeletedAt{
			Time:  *transfer.DeletedAt,
			Valid: true,
		}
	}

	return model
}
