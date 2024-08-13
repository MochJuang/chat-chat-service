package model

type CreateMessageRequest struct {
	SenderId       uint   `json:"sender_id" validate:"required"`
	Content        string `json:"content" validate:"required"`
	ConversationId uint   `validate:"required"`
}
