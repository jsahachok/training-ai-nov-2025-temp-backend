package controllers

import (
	"strconv"
	"workshop4-backend/domain/entities"
	"workshop4-backend/usecases"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TransferController struct {
	transferUseCase *usecases.TransferUseCase
	validator       *validator.Validate
}

func NewTransferController(transferUseCase *usecases.TransferUseCase) *TransferController {
	return &TransferController{
		transferUseCase: transferUseCase,
		validator:       validator.New(),
	}
}

// TransferPoints handles POST /users/{id}/transfer
func (ctrl *TransferController) TransferPoints(c *fiber.Ctx) error {
	fromUserIDParam := c.Params("id")

	fromUserID, err := strconv.ParseUint(fromUserIDParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	var req entities.TransferRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Cannot parse JSON",
		})
	}

	// Validate input
	if err := ctrl.validator.Struct(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	transfer, err := ctrl.transferUseCase.TransferPoints(uint(fromUserID), &req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Transfer completed successfully",
		"data": entities.TransferResponse{
			ID:            transfer.ID,
			FromUserID:    transfer.FromUserID,
			ToUserID:      transfer.ToUserID,
			Amount:        transfer.Amount,
			Description:   transfer.Description,
			Status:        transfer.Status,
			TransferredAt: transfer.TransferredAt,
			CreatedAt:     transfer.CreatedAt,
		},
	})
}

// GetTransferHistory handles GET /users/{id}/transfer
func (ctrl *TransferController) GetTransferHistory(c *fiber.Ctx) error {
	userIDParam := c.Params("id")

	userID, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	// Get pagination parameters
	limitParam := c.Query("limit", "10")
	offsetParam := c.Query("offset", "0")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100 // Max limit
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil || offset < 0 {
		offset = 0
	}

	history, err := ctrl.transferUseCase.GetTransferHistory(uint(userID), limit, offset)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    history.Transfers,
		"total":   history.Total,
		"limit":   limit,
		"offset":  offset,
	})
}

// GetTransferByID handles GET /transfer/{id}
func (ctrl *TransferController) GetTransferByID(c *fiber.Ctx) error {
	transferIDParam := c.Params("id")

	transferID, err := strconv.ParseUint(transferIDParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid transfer ID",
		})
	}

	transfer, err := ctrl.transferUseCase.GetTransferByID(uint(transferID))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": entities.TransferResponse{
			ID:            transfer.ID,
			FromUserID:    transfer.FromUserID,
			ToUserID:      transfer.ToUserID,
			Amount:        transfer.Amount,
			Description:   transfer.Description,
			Status:        transfer.Status,
			TransferredAt: transfer.TransferredAt,
			CreatedAt:     transfer.CreatedAt,
		},
	})
}
