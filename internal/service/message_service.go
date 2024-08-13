package service

import (
	"chat-service/internal/entity"
	e "chat-service/internal/exception"
	"chat-service/internal/model"
	"chat-service/internal/repository"
	"chat-service/internal/utils"
	"time"
)

type MessageService interface {
	CreateMessage(messageDTO *model.CreateMessageRequest) (*model.MessageResponse, error)
	GetMessages(conversationID int64) ([]*model.MessageResponse, error)
}

type messageService struct {
	conversationRepo repository.ConversationRepository
	messageRepo      repository.MessageRepository
}

func NewMessageService(messageRepo repository.MessageRepository, conversationRepo repository.ConversationRepository) MessageService {
	return &messageService{messageRepo: messageRepo, conversationRepo: conversationRepo}
}

func (s *messageService) CreateMessage(request *model.CreateMessageRequest) (*model.MessageResponse, error) {

	err := utils.Validate(request)
	if err != nil {
		return nil, err
	}

	var conversation *entity.Conversation
	conversation, err = s.conversationRepo.GetConversationByID(request.ConversationId)
	if err != nil {
		return nil, e.NotFound("Conversation not found")
	}

	message := &entity.Message{
		ConversationID: conversation.ID,
		SenderID:       request.SenderId,
		Content:        request.Content,
		SendAt:         time.Now(),
	}
	err = s.messageRepo.SaveMessage(message)
	if err != nil {
		return nil, e.Internal(err)
	}

	return model.ToMessageResponse(message), nil
}

func (s *messageService) GetMessages(conversationID int64) ([]*model.MessageResponse, error) {
	messages, err := s.messageRepo.GetMessagesByConversationID(conversationID)
	if err != nil {
		return nil, e.NotFound("Messages not found")
	}

	var messageResponses []*model.MessageResponse
	for _, message := range messages {
		messageResponses = append(messageResponses, model.ToMessageResponse(message))
	}
	return messageResponses, nil
}
