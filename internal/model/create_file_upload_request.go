package model

type CreateFileUploadRequest struct {
	UserID int    `json:"user_id" validate:"required"`
	File   string `json:"file" validate:"required"`
}
