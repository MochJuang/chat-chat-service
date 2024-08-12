package model

import (
	"chat-service/internal/entity"
	"time"
)

type MessageResponse struct {
	ID             uint      `json:"id"`
	ConversationID uint      `json:"conversation_id"`
	SenderID       uint      `json:"sender_id"`
	Content        string    `json:"content"`
	SentAt         time.Time `json:"sent_at"`
}

func ToMessageResponse(message *entity.Message) *MessageResponse {
	return &MessageResponse{
		ID:             message.ID,
		ConversationID: message.ConversationID,
		SenderID:       message.SenderID,
		Content:        message.Content,
		SentAt:         message.SendAt,
	}
}
