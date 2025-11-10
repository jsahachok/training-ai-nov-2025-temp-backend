package entities

import "time"

type Transfer struct {
	ID          uint       `json:"id"`
	FromUserID  uint       `json:"fromUserId"`
	ToUserID    uint       `json:"toUserId"`
	Points      int        `json:"points"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status"` // pending, completed, failed
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

type TransferRequest struct {
	ToUserID    uint   `json:"toUserId" validate:"required"`
	Points      int    `json:"points" validate:"required,min=1"`
	Description string `json:"description,omitempty"`
}

type TransferResponse struct {
	ID          uint      `json:"id"`
	FromUserID  uint      `json:"fromUserId"`
	ToUserID    uint      `json:"toUserId"`
	Points      int       `json:"points"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

type TransferHistoryResponse struct {
	Transfers []TransferResponse `json:"transfers"`
	Total     int                `json:"total"`
}
