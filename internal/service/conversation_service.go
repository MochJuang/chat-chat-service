package service

import (
	"chat-service/internal/entity"
	e "chat-service/internal/exception"
	"chat-service/internal/model"
	"chat-service/internal/repository"
	"chat-service/internal/utils"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ConversationService interface {
	CreateConversation(request *model.CreateConversationRequest) (*model.ConversationResponse, error)
	GetConversationByUuid(conversationUuid string) (*model.ConversationResponse, error)
	GetAllConversations() ([]*model.ConversationResponse, error)
}

type conversationService struct {
	conversationRepo repository.ConversationRepository
}

func NewConversationService(repo repository.ConversationRepository) ConversationService {
	return &conversationService{conversationRepo: repo}
}

func (s *conversationService) CreateConversation(request *model.CreateConversationRequest) (*model.ConversationResponse, error) {
	err := utils.Validate(request)
	if err != nil {
		return nil, err
	}
	cvsUuid := uuid.New().String()
	conversation := &entity.Conversation{
		UUID:      cvsUuid,
		CreatedAt: time.Now(),
		Type:      entity.ConversationTypePrivate,
	}

	var userIds []uint
	for _, userID := range request.Participants {
		user, err := s.conversationRepo.GetUserByUuid(userID)
		if err != nil {
			return nil, e.NotFound("user not found")
		}

		conversation.Participants = append(conversation.Participants, *user)
		userIds = append(userIds, user.ID)
	}

	err = s.conversationRepo.CheckExistingConversation(conversation)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.Internal(err)
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = s.conversationRepo.CreateConversation(conversation)
		if err != nil {
			return nil, e.Internal(err)
		}
	}

	conversation, err = s.conversationRepo.GetConversationByUuid(conversation.UUID)
	if err != nil {
		return nil, e.Internal(err)
	}

	response := model.ToConversationResponse(conversation)
	return response, nil
}

func (s *conversationService) GetConversationByUuid(conversationUuid string) (*model.ConversationResponse, error) {
	conversation, err := s.conversationRepo.GetConversationByUuid(conversationUuid)
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
