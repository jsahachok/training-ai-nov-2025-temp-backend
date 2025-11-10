package infrastructure

import (
	"time"
	"workshop4-backend/domain/entities"

	"gorm.io/gorm"
)

type UserModel struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	FirstName   string `gorm:"column:first_name;type:varchar(100);not null"`
	LastName    string `gorm:"column:last_name;type:varchar(100);not null"`
	Email       string `gorm:"column:email;type:varchar(255);uniqueIndex;not null"`
	Phone       string `gorm:"column:phone;type:varchar(20);not null"`
	DateOfBirth string `gorm:"column:date_of_birth;type:date;not null"`
	Address     string `gorm:"column:address;type:text;not null"`
	City        string `gorm:"column:city;type:varchar(100);not null"`
	Country     string `gorm:"column:country;type:varchar(100);not null"`
	PostalCode  string `gorm:"column:postal_code;type:varchar(20);not null"`
	Avatar      string `gorm:"column:avatar;type:varchar(500);default:''"`
	Points      float64 `gorm:"column:points;type:decimal(10,2);default:0.00;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
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
		Points:      u.Points,
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
		Points:      user.Points,
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
