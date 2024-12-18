package postgresql

import (
	"chat-service/internal/entity"
	"chat-service/internal/repository"
	"errors"
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

func (r *conversationRepository) GetConversationByUuid(conversationUuid string) (*entity.Conversation, error) {
	var conversation entity.Conversation
	err := r.db.Preload("Participants").Preload("Messages").Where("uuid = ?", conversationUuid).First(&conversation).Error
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

func (r *conversationRepository) GetUserByUuid(uuid string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *conversationRepository) GetListConversationByUserUuid(userUuid string) ([]*entity.Conversation, error) {
	var conversations []*entity.Conversation
	err := r.db.Preload("Participants").
		Preload("Messages").
		Joins("JOIN conversation_participants ON conversation_participants.conversation_uuid = conversations.uuid").
		Where("conversation_participants.user_uuid = ?", userUuid).
		Find(&conversations).Error
	if err != nil {
		return nil, err
	}
	return conversations, nil
}

func (r *conversationRepository) CheckExistingConversation(conversation *entity.Conversation) error {

	if len(conversation.Participants) != 2 {
		return errors.New("A conversation must have exactly two participants.")
	}

	var existingConversationUuid string
	err := r.db.Raw(`
		SELECT c.uuid
		FROM conversation_participants a
		JOIN conversation_participants b ON a.conversation_id = b.conversation_id
		JOIN conversations c ON a.conversation_id = c.id
		WHERE a.user_id = ? AND b.user_id = ?
	`, conversation.Participants[0].ID, conversation.Participants[1].ID).Scan(&existingConversationUuid).Error

	if err != nil {
		return err
	}
	
	if existingConversationUuid == "" {
		return gorm.ErrRecordNotFound
	}

	conversation.UUID = existingConversationUuid

	return err
}
