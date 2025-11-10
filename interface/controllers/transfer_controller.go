package controllers

import (
	"strconv"
	"workshop4-backend/domain/entities"
	"workshop4-backend/internal/constants"
	"workshop4-backend/internal/utils"
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
		return utils.BadRequestResponse(c, "Invalid user ID format")
	}

	var req entities.TransferRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "Cannot parse JSON body")
	}

	// Validate input
	if err := ctrl.validator.Struct(&req); err != nil {
		return utils.ValidationErrorResponse(c, err.Error())
	}

	transfer, err := ctrl.transferUseCase.TransferPoints(uint(fromUserID), &req)
	if err != nil {
		switch err.Error() {
		case "invalid from user ID":
			return utils.BadRequestResponse(c, err.Error())
		case "cannot transfer points to yourself":
			return utils.ErrorResponse(c, fiber.StatusBadRequest, constants.ErrCodeSelfTransferNotAllowed, err.Error())
		case "from user not found", "to user not found":
			return utils.NotFoundResponse(c, "User")
		case "insufficient points":
			return utils.ErrorResponse(c, fiber.StatusBadRequest, constants.ErrCodeInsufficientBalance, err.Error())
		default:
			return utils.ErrorResponse(c, fiber.StatusBadRequest, constants.ErrCodeTransferFailed, err.Error())
		}
	}

	transferResponse := entities.TransferResponse{
		ID:            transfer.ID,
		FromUserID:    transfer.FromUserID,
		ToUserID:      transfer.ToUserID,
		Amount:        transfer.Amount,
		Description:   transfer.Description,
		Status:        transfer.Status,
		TransferredAt: transfer.TransferredAt,
		CreatedAt:     transfer.CreatedAt,
	}

	return utils.CreatedResponse(c, transferResponse, "Transfer completed successfully")
}

// GetTransferHistory handles GET /users/{id}/transfer
func (ctrl *TransferController) GetTransferHistory(c *fiber.Ctx) error {
	userIDParam := c.Params("id")
	
	userID, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid user ID format")
	}

	// Get pagination parameters with defaults
	limitParam := c.Query("limit", strconv.Itoa(constants.DefaultLimit))
	offsetParam := c.Query("offset", strconv.Itoa(constants.DefaultOffset))

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = constants.DefaultLimit
	}
	if limit > constants.MaxLimit {
		limit = constants.MaxLimit
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil || offset < 0 {
		offset = constants.DefaultOffset
	}

	history, err := ctrl.transferUseCase.GetTransferHistory(uint(userID), limit, offset)
	if err != nil {
		if err.Error() == "invalid user ID" {
			return utils.BadRequestResponse(c, err.Error())
		}
		if err.Error() == "user not found" {
			return utils.NotFoundResponse(c, "User")
		}
		return utils.InternalErrorResponse(c, err.Error())
	}

	return utils.PaginatedResponse(c, history.Transfers, history.Total, limit, offset)
}

// GetTransferByID handles GET /transfer/{id}
func (ctrl *TransferController) GetTransferByID(c *fiber.Ctx) error {
	transferIDParam := c.Params("id")
	
	transferID, err := strconv.ParseUint(transferIDParam, 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid transfer ID format")
	}

	transfer, err := ctrl.transferUseCase.GetTransferByID(uint(transferID))
	if err != nil {
		if err.Error() == "invalid transfer ID" {
			return utils.BadRequestResponse(c, err.Error())
		}
		return utils.NotFoundResponse(c, "Transfer")
	}

	transferResponse := entities.TransferResponse{
		ID:            transfer.ID,
		FromUserID:    transfer.FromUserID,
		ToUserID:      transfer.ToUserID,
		Amount:        transfer.Amount,
		Description:   transfer.Description,
		Status:        transfer.Status,
		TransferredAt: transfer.TransferredAt,
		CreatedAt:     transfer.CreatedAt,
	}

	return utils.SuccessResponse(c, transferResponse)
}
