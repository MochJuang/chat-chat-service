package http

import (
	"chat-service/internal/model"
	"chat-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type MessageController struct {
	MessageService service.MessageService
}

func NewMessageController(service service.MessageService) *MessageController {
	return &MessageController{MessageService: service}
}

func (h *MessageController) CreateMessage(c *fiber.Ctx) error {
	messageDTO := new(model.CreateMessageRequest)
	if err := c.BodyParser(messageDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	message, err := h.MessageService.CreateMessage(messageDTO)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(message)
}

func (h *MessageController) GetMessages(c *fiber.Ctx) error {
	conversationID, err := c.ParamsInt("conversationID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid conversation ID"})
	}

	messages, err := h.MessageService.GetMessages(int64(conversationID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(messages)
}
