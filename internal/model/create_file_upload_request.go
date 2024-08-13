package model

type CreateFileUploadRequest struct {
	UserID uint   `json:"user_id" validate:"required"`
	File   string `json:"file" validate:"required"`
}
