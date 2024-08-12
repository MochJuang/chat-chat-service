package repository

import "chat-service/internal/entity"

type MessageRepository interface {
	SaveMessage(msg *entity.Message) error
	GetMessagesByConversationID(conversationID int64) ([]*entity.Message, error)
}
