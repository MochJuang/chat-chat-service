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
	conversationRepo := postgresql.NewConversationRepository(cfg.DB)
	fileUploadRepo := postgresql.NewFileUploadRepository(cfg.DB)

	// Services
	messageService := service.NewMessageService(messageRepo, conversationRepo)
	conversationService := service.NewConversationService(conversationRepo)
	fileUploadService := service.NewFileUploadService(fileUploadRepo)

	// Handlers
	messageController := httpdelivery.NewMessageController(messageService)
	conversationController := httpdelivery.NewConversationHandler(conversationService)
	fileUploadController := httpdelivery.NewFileUploadController(fileUploadService)

	// Public routes
	app.Post("/conversations/:id/messages", messageController.CreateMessage)
	app.Get("/conversations/:id/messages/", messageController.GetMessages)

	app.Post("/conversations", conversationController.CreateConversation)
	app.Get("/conversations/:conversationID", conversationController.GetConversationByID)
	app.Get("/conversations", conversationController.GetAllConversations)

	app.Post("/files", fileUploadController.UploadFile)
	app.Get("/files/:fileID", fileUploadController.GetFileByID)

}
