package service

import (
	"chat-service/internal/entity"
	e "chat-service/internal/exception"
	"chat-service/internal/model"
	"chat-service/internal/repository"
	"time"
)

type MessageService interface {
	CreateMessage(messageDTO *model.CreateMessageRequest) (*model.MessageResponse, error)
	GetMessages(conversationID int64) ([]*model.MessageResponse, error)
}

type messageService struct {
	messageRepo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{messageRepo: repo}
}

func (s *messageService) CreateMessage(request *model.CreateMessageRequest) (*model.MessageResponse, error) {
	var conversation entity.Conversation

	message := &entity.Message{
		ConversationID: conversation.ID,
		SenderID:       request.UserID,
		Content:        request.Content,
		SendAt:         time.Now(),
	}
	err := s.messageRepo.SaveMessage(message)
	if err != nil {
		return nil, e.Internal(err)
	}

	response := model.ToMessageResponse(message)
	return response, nil
}

func (s *messageService) GetMessages(conversationID int64) ([]*model.MessageResponse, error) {
	messages, err := s.messageRepo.GetMessagesByConversationID(conversationID)
	if err != nil {
		return nil, e.NotFound("Messages not found")
	}

	var response []*model.MessageResponse
	for _, message := range messages {
		response = append(response, model.ToMessageResponse(message))
	}
	return response, nil
}
