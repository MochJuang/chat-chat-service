package model

import "chat-service/internal/entity"

type ConversationResponse struct {
	UUID         string   `json:"uuid"`
	Participants []string `json:"participants"`
	CreatedAt    string   `json:"created_at"`
}

func ToConversationResponse(conversation *entity.Conversation) *ConversationResponse {
	participantIds := make([]string, len(conversation.Participants))
	for i, participant := range conversation.Participants {
		participantIds[i] = participant.UUID

	}

	return &ConversationResponse{
		UUID:         conversation.UUID,
		Participants: participantIds,
		CreatedAt:    conversation.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
