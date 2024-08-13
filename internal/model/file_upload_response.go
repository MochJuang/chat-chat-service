package model

import (
	"chat-service/internal/entity"
	"time"
)

type FileUploadResponse struct {
	ID           int                   `json:"id"`
	UserID       int                   `json:"user_id"`
	FileURL      string                `json:"file_url"`
	Conversation *ConversationResponse `json:"conversation,omitempty"`
	UploadedAt   time.Time             `json:"uploaded_at"`
}

func ToFileUploadResponse(fileUpload *entity.FileUpload, conversation *ConversationResponse) *FileUploadResponse {
	return &FileUploadResponse{
		ID:           int(fileUpload.ID),
		UserID:       int(fileUpload.UserID),
		FileURL:      fileUpload.FileURL,
		Conversation: conversation,
		UploadedAt:   fileUpload.CreatedAt,
	}
}
