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

type UserController struct {
	userUseCase *usecases.UserUseCase
	validator   *validator.Validate
}

func NewUserController(userUseCase *usecases.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
		validator:   validator.New(),
	}
}

// GetAllUsers handles GET /users
func (ctrl *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := ctrl.userUseCase.GetAllUsers()
	if err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}

	return utils.SuccessResponse(c, users)
}

// GetUserByID handles GET /users/:id
func (ctrl *UserController) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid user ID format")
	}

	user, err := ctrl.userUseCase.GetUserByID(uint(id))
	if err != nil {
		if err.Error() == "invalid user ID" {
			return utils.BadRequestResponse(c, err.Error())
		}
		return utils.NotFoundResponse(c, "User")
	}

	return utils.SuccessResponse(c, user)
}

// CreateUser handles POST /users
func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	var req entities.UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "Cannot parse JSON body")
	}

	// Validate input
	if err := ctrl.validator.Struct(&req); err != nil {
		return utils.ValidationErrorResponse(c, err.Error())
	}

	user, err := ctrl.userUseCase.CreateUser(&req)
	if err != nil {
		if err.Error() == "user with this email already exists" {
			return utils.ErrorResponse(c, fiber.StatusConflict, constants.ErrCodeUserEmailExists, err.Error())
		}
		return utils.InternalErrorResponse(c, err.Error())
	}

	return utils.CreatedResponse(c, user, "User created successfully")
}

// UpdateUser handles PUT /users/:id
func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid user ID format")
	}

	var req entities.UserUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "Cannot parse JSON body")
	}

	// Validate input
	if err := ctrl.validator.Struct(&req); err != nil {
		return utils.ValidationErrorResponse(c, err.Error())
	}

	user, err := ctrl.userUseCase.UpdateUser(uint(id), &req)
	if err != nil {
		if err.Error() == "invalid user ID" {
			return utils.BadRequestResponse(c, err.Error())
		}
		if err.Error() == "user not found" {
			return utils.NotFoundResponse(c, "User")
		}
		if err.Error() == "user with this email already exists" {
			return utils.ErrorResponse(c, fiber.StatusConflict, constants.ErrCodeUserEmailExists, err.Error())
		}
		return utils.InternalErrorResponse(c, err.Error())
	}

	return utils.SuccessResponse(c, user, constants.MsgUpdated)
}

// DeleteUser handles DELETE /users/:id
func (ctrl *UserController) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid user ID format")
	}

	err = ctrl.userUseCase.DeleteUser(uint(id))
	if err != nil {
		if err.Error() == "invalid user ID" {
			return utils.BadRequestResponse(c, err.Error())
		}
		if err.Error() == "user not found" {
			return utils.NotFoundResponse(c, "User")
		}
		return utils.InternalErrorResponse(c, err.Error())
	}

	return utils.SuccessResponse(c, nil, constants.MsgDeleted)
}
