package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	FirstName   string         `json:"firstName" validate:"required" gorm:"not null"`
	LastName    string         `json:"lastName" validate:"required" gorm:"not null"`
	Email       string         `json:"email" validate:"required,email" gorm:"uniqueIndex;not null"`
	Phone       string         `json:"phone" validate:"required" gorm:"not null"`
	DateOfBirth string         `json:"dateOfBirth" validate:"required" gorm:"not null"`
	Address     string         `json:"address" validate:"required" gorm:"not null"`
	City        string         `json:"city" validate:"required" gorm:"not null"`
	Country     string         `json:"country" validate:"required" gorm:"not null"`
	PostalCode  string         `json:"postalCode" validate:"required" gorm:"not null"`
	Avatar      string         `json:"avatar" gorm:"default:''"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
