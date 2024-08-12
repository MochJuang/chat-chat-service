package service

import (
	"chat-service/internal/entity"
	e "chat-service/internal/exception"
	"chat-service/internal/model"
	"chat-service/internal/repository"
	"time"
)

type ConversationService interface {
	CreateConversation(request *model.CreateConversationRequest) (*model.ConversationResponse, error)
	GetConversationByID(conversationID uint) (*model.ConversationResponse, error)
	GetAllConversations() ([]*model.ConversationResponse, error)
}

type conversationService struct {
	conversationRepo repository.ConversationRepository
}

func NewConversationService(repo repository.ConversationRepository) ConversationService {
	return &conversationService{conversationRepo: repo}
}

func (s *conversationService) CreateConversation(request *model.CreateConversationRequest) (*model.ConversationResponse, error) {
	conversation := &entity.Conversation{
		CreatedAt: time.Now(),
	}

	for _, userID := range request.Participants {
		user := entity.User{ID: userID}
		conversation.Participants = append(conversation.Participants, user)
	}

	err := s.conversationRepo.CreateConversation(conversation)
	if err != nil {
		return nil, e.Internal(err)
	}

	response := model.ToConversationResponse(conversation)
	return response, nil
}

func (s *conversationService) GetConversationByID(conversationID uint) (*model.ConversationResponse, error) {
	conversation, err := s.conversationRepo.GetConversationByID(conversationID)
	if err != nil {
		return nil, e.NotFound("conversation not found")
	}

	response := model.ToConversationResponse(conversation)
	return response, nil
}

func (s *conversationService) GetAllConversations() ([]*model.ConversationResponse, error) {
	conversations, err := s.conversationRepo.GetAllConversations()
	if err != nil {
		return nil, e.NotFound("conversation not found")
	}

	var response []*model.ConversationResponse
	for _, conversation := range conversations {
		response = append(response, model.ToConversationResponse(conversation))
	}
	return response, nil
}
