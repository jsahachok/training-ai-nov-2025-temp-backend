package infrastructure

import (
	"time"
	"workshop4-backend/domain/entities"

	"gorm.io/gorm"
)

type TransferModel struct {
	ID            uint           `gorm:"primaryKey;autoIncrement"`
	FromUserID    uint           `gorm:"column:from_user_id;not null;index"`
	ToUserID      uint           `gorm:"column:to_user_id;not null;index"`
	Amount        float64        `gorm:"column:amount;not null;type:decimal(10,2)"`
	Description   string         `gorm:"column:description;type:varchar(255);default:''"`
	Status        string         `gorm:"column:status;type:varchar(50);not null;default:'pending'"`
	TransferredAt time.Time      `gorm:"column:transferred_at;not null"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index"`

	// Foreign key relationships
	FromUser UserModel `gorm:"foreignKey:FromUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ToUser   UserModel `gorm:"foreignKey:ToUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
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
		ID:            t.ID,
		FromUserID:    t.FromUserID,
		ToUserID:      t.ToUserID,
		Amount:        t.Amount,
		Description:   t.Description,
		Status:        t.Status,
		TransferredAt: t.TransferredAt,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
		DeletedAt:     deletedAt,
	}
}

func FromTransferEntity(transfer *entities.Transfer) *TransferModel {
	model := &TransferModel{
		ID:            transfer.ID,
		FromUserID:    transfer.FromUserID,
		ToUserID:      transfer.ToUserID,
		Amount:        transfer.Amount,
		Description:   transfer.Description,
		Status:        transfer.Status,
		TransferredAt: transfer.TransferredAt,
		CreatedAt:     transfer.CreatedAt,
		UpdatedAt:     transfer.UpdatedAt,
	}

	if transfer.DeletedAt != nil {
		model.DeletedAt = gorm.DeletedAt{
			Time:  *transfer.DeletedAt,
			Valid: true,
		}
	}

	return model
}
