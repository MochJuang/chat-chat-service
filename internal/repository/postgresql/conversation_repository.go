package postgresql

import (
	"chat-service/internal/entity"
	"chat-service/internal/repository"
	"gorm.io/gorm"
)

type conversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) repository.ConversationRepository {
	return &conversationRepository{db: db}
}

func (r *conversationRepository) CreateConversation(conversation *entity.Conversation) error {
	return r.db.Create(conversation).Error
}

func (r *conversationRepository) GetConversationByID(conversationID uint) (*entity.Conversation, error) {
	var conversation entity.Conversation
	err := r.db.Preload("Participants").Preload("Messages").First(&conversation, conversationID).Error
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *conversationRepository) GetAllConversations() ([]*entity.Conversation, error) {
	var conversations []*entity.Conversation
	err := r.db.Preload("Participants").Preload("Messages").Find(&conversations).Error
	if err != nil {
		return nil, err
	}
	return conversations, nil
}
