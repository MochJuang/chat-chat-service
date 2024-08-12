package route

import (
	"chat-service/internal/config"
	httpdelivery "chat-service/internal/delivery/http"
	middleware "chat-service/internal/delivery/http/midlleware"
	"chat-service/internal/repository/postgresql"
	"chat-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg config.Config) {
	// Initialize http
	app.Use(middleware.ErrorHandlerMiddleware)

	messageRepo := postgresql.NewMessageRepository(cfg.DB)

	// Services
	messageService := service.NewMessageService(messageRepo)

	// Handlers
	messageController := httpdelivery.NewMessageController(messageService)

	// Public routes
	app.Post("/messages", messageController.CreateMessage)
	app.Get("/messages/:conversationID", messageController.GetMessages)

}
