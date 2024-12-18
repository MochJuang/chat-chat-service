package http

import (
	e "chat-service/internal/exception"
	"chat-service/internal/model"
	"chat-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ConversationHandler struct {
	ConversationService service.ConversationService
}

func NewConversationHandler(service service.ConversationService) *ConversationHandler {
	return &ConversationHandler{ConversationService: service}
}

func (h *ConversationHandler) CreateConversation(c *fiber.Ctx) error {
	conversationDTO := new(model.CreateConversationRequest)
	if err := c.BodyParser(conversationDTO); err != nil {
		return e.BadRequest(err)
	}

	conversation, err := h.ConversationService.CreateConversation(conversationDTO)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(conversation)
}

func (h *ConversationHandler) GetConversationByUuid(c *fiber.Ctx) error {
	conversationID := c.Params("conversationID")

	conversation, err := h.ConversationService.GetConversationByUuid(conversationID)
	if err != nil {
		return err
	}

	return c.JSON(conversation)
}

func (h *ConversationHandler) GetAllConversations(c *fiber.Ctx) error {
	conversations, err := h.ConversationService.GetAllConversations()
	if err != nil {
		return err
	}

	return c.JSON(conversations)
}
