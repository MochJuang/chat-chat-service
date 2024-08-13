package repository

import "chat-service/internal/entity"

type FileUploadRepository interface {
	SaveFileUpload(file *entity.FileUpload) error
	GetFilesByMessageID(messageID uint) ([]*entity.FileUpload, error)
	GetFilesByID(id uint) (*entity.FileUpload, error)
}
