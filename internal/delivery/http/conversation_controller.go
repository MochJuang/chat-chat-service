package http

import (
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	conversation, err := h.ConversationService.CreateConversation(conversationDTO)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(conversation)
}

func (h *ConversationHandler) GetConversationByID(c *fiber.Ctx) error {
	conversationID, err := c.ParamsInt("conversationID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid conversation ID"})
	}

	conversation, err := h.ConversationService.GetConversationByID(uint(conversationID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(conversation)
}

func (h *ConversationHandler) GetAllConversations(c *fiber.Ctx) error {
	conversations, err := h.ConversationService.GetAllConversations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(conversations)
}
