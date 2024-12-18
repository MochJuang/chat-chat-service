package model

type CreateMessageRequest struct {
	SenderUuid       string `json:"sender_id" validate:"required"`
	Content          string `json:"content" validate:"required"`
	ConversationUuid string `validate:"required"`
}
