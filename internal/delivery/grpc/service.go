package grpc

import (
	"chat-service/internal/model"
	"chat-service/internal/service"
	"context"
	"github.com/MochJuang/chat-grpc/service/chat"
)

type ChatService struct {
	chat.UnimplementedChatServiceServer
	conversationService service.ConversationService
	messageService      service.MessageService
}

func NewChatService(conversationService service.ConversationService, messageService service.MessageService) chat.ChatServiceServer {
	return &ChatService{
		conversationService: conversationService,
		messageService:      messageService,
	}
}

func (s *ChatService) AddMessageToConversation(ctx context.Context, req *chat.AddMessageRequest) (*chat.AddMessageResponse, error) {

	conversation, err := s.conversationService.GetConversationByID(uint(req.ConversationId))
	if err != nil {
		return nil, err
	}

	isAuthorized := false
	for _, participant := range conversation.Participants {
		if uint(req.SenderId) == participant {
			isAuthorized = true
			break
		}
	}
	if !isAuthorized {
		return &chat.AddMessageResponse{
			Success: false,
			Message: "You are not authorized to send message to this conversation",
		}, nil
	}

	request := &model.CreateMessageRequest{
		ConversationId: uint(req.ConversationId),
		SenderId:       uint(req.SenderId),
		Content:        req.Content,
	}

	if _, err = s.messageService.CreateMessage(request); err != nil {
		return &chat.AddMessageResponse{
			Success: false,
			Message: "Failed to add message",
		}, err
	}

	return &chat.AddMessageResponse{
		Success: true,
		Message: "Message added successfully",
	}, nil
}

func (s *ChatService) GetConversationDetails(ctx context.Context, req *chat.ConversationRequest) (*chat.ConversationResponse, error) {
	conversation, err := s.conversationService.GetConversationByID(uint(req.ConversationId))
	if err != nil {
		return nil, err
	}

	response := &chat.ConversationResponse{
		Id:        uint32(conversation.ID),
		CreatedAt: conversation.CreatedAt,
	}
	for _, participant := range conversation.Participants {
		response.ParticipantIds = append(response.ParticipantIds, uint32(participant))
	}

	return response, nil
}
