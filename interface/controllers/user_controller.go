package controllers

import (
	"strconv"
	"workshop4-backend/domain/entities"
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
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch users",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    users,
		"total":   len(users),
	})
}

// GetUserByID handles GET /users/{id}
func (ctrl *UserController) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	user, err := ctrl.userUseCase.GetUserByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

// CreateUser handles POST /users
func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	var req entities.UserCreateRequest

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

	user, err := ctrl.userUseCase.CreateUser(&req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"data":    user,
	})
}

// UpdateUser handles PUT /users/{id}
func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	var req entities.UserUpdateRequest
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

	user, err := ctrl.userUseCase.UpdateUser(uint(id), &req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
		"data":    user,
	})
}

// DeleteUser handles DELETE /users/{id}
func (ctrl *UserController) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	err = ctrl.userUseCase.DeleteUser(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})
}
