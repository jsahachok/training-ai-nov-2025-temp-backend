package entities

import "time"

type User struct {
	ID          uint       `json:"id"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	DateOfBirth string     `json:"dateOfBirth"`
	Address     string     `json:"address"`
	City        string     `json:"city"`
	Country     string     `json:"country"`
	PostalCode  string     `json:"postalCode"`
	Avatar      string     `json:"avatar"`
	Points      int        `json:"points"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

type UserCreateRequest struct {
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"required"`
	DateOfBirth string `json:"dateOfBirth" validate:"required"`
	Address     string `json:"address" validate:"required"`
	City        string `json:"city" validate:"required"`
	Country     string `json:"country" validate:"required"`
	PostalCode  string `json:"postalCode" validate:"required"`
	Avatar      string `json:"avatar"`
}

type UserUpdateRequest struct {
	FirstName   *string `json:"firstName,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	Email       *string `json:"email,omitempty" validate:"omitempty,email"`
	Phone       *string `json:"phone,omitempty"`
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	Address     *string `json:"address,omitempty"`
	City        *string `json:"city,omitempty"`
	Country     *string `json:"country,omitempty"`
	PostalCode  *string `json:"postalCode,omitempty"`
	Avatar      *string `json:"avatar,omitempty"`
}
