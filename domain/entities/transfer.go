package entities

import "time"

type Transfer struct {
	ID            uint       `json:"id"`
	FromUserID    uint       `json:"fromUserId"`
	ToUserID      uint       `json:"toUserId"`
	Amount        float64    `json:"amount"`
	Description   string     `json:"description,omitempty"`
	Status        string     `json:"status"` // pending, completed, failed
	TransferredAt time.Time  `json:"transferredAt"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty"`
}

type TransferRequest struct {
	ToUserID    uint    `json:"toUserId" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,min=0.01"`
	Description string  `json:"description,omitempty"`
}

type TransferResponse struct {
	ID            uint      `json:"id"`
	FromUserID    uint      `json:"fromUserId"`
	ToUserID      uint      `json:"toUserId"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description,omitempty"`
	Status        string    `json:"status"`
	TransferredAt time.Time `json:"transferredAt"`
	CreatedAt     time.Time `json:"createdAt"`
}

type TransferHistoryResponse struct {
	Transfers []TransferResponse `json:"transfers"`
	Total     int                `json:"total"`
}
