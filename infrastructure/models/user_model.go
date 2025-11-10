package infrastructure

import (
	"time"
	"workshop4-backend/domain/entities"

	"gorm.io/gorm"
)

type UserModel struct {
	ID          uint   `gorm:"primaryKey"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Email       string `gorm:"uniqueIndex;not null"`
	Phone       string `gorm:"not null"`
	DateOfBirth string `gorm:"not null"`
	Address     string `gorm:"not null"`
	City        string `gorm:"not null"`
	Country     string `gorm:"not null"`
	PostalCode  string `gorm:"not null"`
	Avatar      string `gorm:"default:''"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}

func (u *UserModel) ToEntity() *entities.User {
	var deletedAt *time.Time
	if u.DeletedAt.Valid {
		deletedAt = &u.DeletedAt.Time
	}

	return &entities.User{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		Phone:       u.Phone,
		DateOfBirth: u.DateOfBirth,
		Address:     u.Address,
		City:        u.City,
		Country:     u.Country,
		PostalCode:  u.PostalCode,
		Avatar:      u.Avatar,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}

func FromEntity(user *entities.User) *UserModel {
	model := &UserModel{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Phone:       user.Phone,
		DateOfBirth: user.DateOfBirth,
		Address:     user.Address,
		City:        user.City,
		Country:     user.Country,
		PostalCode:  user.PostalCode,
		Avatar:      user.Avatar,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	if user.DeletedAt != nil {
		model.DeletedAt = gorm.DeletedAt{
			Time:  *user.DeletedAt,
			Valid: true,
		}
	}

	return model
}
