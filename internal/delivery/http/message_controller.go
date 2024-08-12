package http

import (
	e "chat-service/internal/exception"
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
	request := new(model.CreateMessageRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	message, err := h.MessageService.CreateMessage(request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(message)
}

func (h *MessageController) GetMessages(c *fiber.Ctx) error {
	conversationID, err := c.ParamsInt("conversationID")
	if err != nil {
		return e.Validation(err)
	}

	messages, err := h.MessageService.GetMessages(int64(conversationID))
	if err != nil {
		return err
	}

	return c.JSON(messages)
}
