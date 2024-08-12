package model

type CreateConversationRequest struct {
	Participants []uint `json:"participants" validate:"required"`
}
