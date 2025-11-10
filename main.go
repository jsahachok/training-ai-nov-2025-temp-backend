package main

import (
	"log"
	"workshop4-backend/infrastructure"
	infrastructureRepo "workshop4-backend/infrastructure/repositories"
	"workshop4-backend/interface/controllers"
	"workshop4-backend/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize database
	database := infrastructure.NewDatabase()

	// Initialize repositories
	userRepo := infrastructureRepo.NewUserRepository(database.GetDB())
	transferRepo := infrastructureRepo.NewTransferRepository(database.GetDB())

	// Initialize use cases
	userUseCase := usecases.NewUserUseCase(userRepo)
	transferUseCase := usecases.NewTransferUseCase(transferRepo, userRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userUseCase)
	transferController := controllers.NewTransferController(transferUseCase)

	// Create new Fiber instance
	app := fiber.New()

	// Add middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello world",
		})
	})

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// User routes
	app.Get("/users", userController.GetAllUsers)
	app.Get("/users/:id", userController.GetUserByID)
	app.Post("/users", userController.CreateUser)
	app.Put("/users/:id", userController.UpdateUser)
	app.Delete("/users/:id", userController.DeleteUser)

	// Transfer routes
	app.Post("/users/:id/transfer", transferController.TransferPoints)
	app.Get("/users/:id/transfer", transferController.GetTransferHistory)
	app.Get("/transfer/:id", transferController.GetTransferByID)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
