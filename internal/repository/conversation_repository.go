package repository

import "chat-service/internal/entity"

type ConversationRepository interface {
	CreateConversation(conversation *entity.Conversation) error
	GetConversationByID(conversationID uint) (*entity.Conversation, error)
	GetAllConversations() ([]*entity.Conversation, error)
}
