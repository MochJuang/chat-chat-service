package repository

import "chat-service/internal/entity"

type ConversationRepository interface {
	CreateConversation(conversation *entity.Conversation) error
	GetConversationByUuid(conversationUuid string) (*entity.Conversation, error)
	GetAllConversations() ([]*entity.Conversation, error)
	GetListConversationByUserUuid(userUuid string) ([]*entity.Conversation, error)
	CheckExistingConversation(conversation *entity.Conversation) error
	GetUserByUuid(senderUuid string) (*entity.User, error)
}
