package postgresql

import (
	"chat-service/internal/entity"
	"chat-service/internal/repository"
	"gorm.io/gorm"
)

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) repository.MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) SaveMessage(msg *entity.Message) error {
	return r.db.Create(msg).Error
}

func (r *messageRepository) GetMessagesByConversationID(conversationID int64) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := r.db.Where("conversation_id = ?", conversationID).Find(&messages).Error
	return messages, err
}
