package model

type CreateMessageRequest struct {
	UserID  uint   `json:"user_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}
