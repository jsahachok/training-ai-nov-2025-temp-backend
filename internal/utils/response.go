package utils

import (
	"workshop4-backend/domain/entities"
	"workshop4-backend/internal/constants"

	"github.com/gofiber/fiber/v2"
)

// ErrorResponse creates standardized error response
func ErrorResponse(c *fiber.Ctx, statusCode int, errCode, message string, details ...string) error {
	response := entities.APIResponse{
		Success: false,
		Error: &entities.APIError{
			Code:    errCode,
			Message: message,
		},
	}

	if len(details) > 0 {
		response.Error.Details = details[0]
	}

	return c.Status(statusCode).JSON(response)
}

// SuccessResponse creates standardized success response
func SuccessResponse(c *fiber.Ctx, data interface{}, message ...string) error {
	response := entities.APIResponse{
		Success: true,
		Data:    data,
	}

	if len(message) > 0 {
		response.Message = message[0]
	} else {
		response.Message = constants.MsgSuccess
	}

	return c.JSON(response)
}

// CreatedResponse creates standardized 201 response
func CreatedResponse(c *fiber.Ctx, data interface{}, message ...string) error {
	response := entities.APIResponse{
		Success: true,
		Data:    data,
	}

	if len(message) > 0 {
		response.Message = message[0]
	} else {
		response.Message = constants.MsgCreated
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// PaginatedResponse creates standardized paginated response
func PaginatedResponse(c *fiber.Ctx, data interface{}, total, limit, offset int) error {
	page := (offset / limit) + 1
	if limit == 0 {
		page = 1
	}

	response := entities.PaginatedResponse{
		Success: true,
		Data:    data,
		Meta: entities.PaginationMeta{
			Total:  total,
			Limit:  limit,
			Offset: offset,
			Page:   page,
		},
	}

	return c.JSON(response)
}

// ValidationErrorResponse creates validation error response
func ValidationErrorResponse(c *fiber.Ctx, details string) error {
	return ErrorResponse(
		c,
		fiber.StatusBadRequest,
		constants.ErrCodeValidation,
		constants.MsgValidationError,
		details,
	)
}

// NotFoundResponse creates not found error response
func NotFoundResponse(c *fiber.Ctx, resource string) error {
	return ErrorResponse(
		c,
		fiber.StatusNotFound,
		constants.ErrCodeNotFound,
		resource+" not found",
	)
}

// BadRequestResponse creates bad request error response
func BadRequestResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(
		c,
		fiber.StatusBadRequest,
		constants.ErrCodeBadRequest,
		message,
	)
}

// InternalErrorResponse creates internal server error response
func InternalErrorResponse(c *fiber.Ctx, details ...string) error {
	return ErrorResponse(
		c,
		fiber.StatusInternalServerError,
		constants.ErrCodeInternalServer,
		constants.MsgInternalError,
		details...,
	)
}