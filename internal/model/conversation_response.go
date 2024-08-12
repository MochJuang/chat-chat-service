package model

import "chat-service/internal/entity"

type ConversationResponse struct {
	ID           uint   `json:"id"`
	Participants []uint `json:"participants"`
	CreatedAt    string `json:"created_at"`
}

func ToConversationResponse(conversation *entity.Conversation) *ConversationResponse {
	participantIds := make([]uint, len(conversation.Participants))
	for i, participant := range conversation.Participants {
		participantIds[i] = participant.ID

	}

	return &ConversationResponse{
		ID:           conversation.ID,
		Participants: participantIds,
		CreatedAt:    conversation.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
