package main

import (
	"log"
	"workshop4-backend/infrastructure"
	infrastructureRepo "workshop4-backend/infrastructure/repositories"
	"workshop4-backend/interface/controllers"
	"workshop4-backend/internal/utils"
	"workshop4-backend/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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

	// Create new Fiber instance with custom error handler
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if e, ok := err.(*fiber.Error); ok {
				return utils.ErrorResponse(c, e.Code, "FIBER_ERROR", e.Message)
			}
			return utils.InternalErrorResponse(c, err.Error())
		},
	})

	// Add middleware
	app.Use(recover.New()) // Recover from panics
	app.Use(requestid.New()) // Add request ID
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} - ${latency}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return utils.SuccessResponse(c, fiber.Map{
			"status":  "ok",
			"service": "workshop4-backend",
			"version": "1.0.0",
		})
	})

	// API version prefix
	api := app.Group("/api/v1")

	// User routes
	api.Get("/users", userController.GetAllUsers)
	api.Get("/users/:id", userController.GetUserByID)
	api.Post("/users", userController.CreateUser)
	api.Put("/users/:id", userController.UpdateUser)
	api.Delete("/users/:id", userController.DeleteUser)

	// Transfer routes
	api.Post("/users/:id/transfer", transferController.TransferPoints)
	api.Get("/users/:id/transfer", transferController.GetTransferHistory)
	api.Get("/transfer/:id", transferController.GetTransferByID)

	// Legacy routes (for backward compatibility)
	app.Get("/users", userController.GetAllUsers)
	app.Get("/users/:id", userController.GetUserByID)
	app.Post("/users", userController.CreateUser)
	app.Put("/users/:id", userController.UpdateUser)
	app.Delete("/users/:id", userController.DeleteUser)
	app.Post("/users/:id/transfer", transferController.TransferPoints)
	app.Get("/users/:id/transfer", transferController.GetTransferHistory)
	app.Get("/transfer/:id", transferController.GetTransferByID)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return utils.NotFoundResponse(c, "Endpoint")
	})

	// Start server
	log.Printf("🚀 Server starting on port 3000...")
	log.Printf("📚 API Documentation available at: http://localhost:3000/api/v1")
	log.Printf("❤️  Health check available at: http://localhost:3000/health")
	log.Fatal(app.Listen(":3000"))
}
