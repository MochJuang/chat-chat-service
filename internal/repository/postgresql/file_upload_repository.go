package postgresql

import (
	"chat-service/internal/entity"
	"chat-service/internal/repository"
	"gorm.io/gorm"
)

type fileUploadRepository struct {
	db *gorm.DB
}

func NewFileUploadRepository(db *gorm.DB) repository.FileUploadRepository {
	return &fileUploadRepository{db: db}
}

func (r *fileUploadRepository) SaveFileUpload(file *entity.FileUpload) error {
	return r.db.Create(file).Error
}

func (r *fileUploadRepository) GetFilesByMessageID(messageID uint) ([]*entity.FileUpload, error) {
	var files []*entity.FileUpload
	err := r.db.Where("message_id = ?", messageID).Find(&files).Error
	return files, err
}

func (r *fileUploadRepository) GetFilesByID(id uint) (*entity.FileUpload, error) {
	var files *entity.FileUpload
	err := r.db.Where("id = ?", id).First(&files).Error
	return files, err
}
