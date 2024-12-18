package model

type CreateConversationRequest struct {
	Participants []string `json:"participants" validate:"required"`
}
