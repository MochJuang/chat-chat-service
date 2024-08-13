package service

import (
	"chat-service/internal/entity"
	e "chat-service/internal/exception"
	"chat-service/internal/model"
	"chat-service/internal/repository"
	"chat-service/internal/utils"
	"time"
)

type FileUploadService interface {
	UploadFile(fileDTO *model.CreateFileUploadRequest) (*model.FileUploadResponse, error)
	GetFileByID(messageID uint) (*model.FileUploadResponse, error)
}

type fileUploadService struct {
	fileUploadRepo repository.FileUploadRepository
}

func NewFileUploadService(repo repository.FileUploadRepository) FileUploadService {
	return &fileUploadService{fileUploadRepo: repo}
}

func (s *fileUploadService) UploadFile(request *model.CreateFileUploadRequest) (*model.FileUploadResponse, error) {
	err := utils.Validate(request)
	if err != nil {
		return nil, err
	}

	file := &entity.FileUpload{
		UserID:    request.UserID,
		FileURL:   request.File,
		CreatedAt: time.Now(),
	}
	err = s.fileUploadRepo.SaveFileUpload(file)
	if err != nil {
		return nil, e.Internal(err)
	}

	response := model.ToFileUploadResponse(file, nil)
	return response, nil
}

func (s *fileUploadService) GetFileByID(fileID uint) (*model.FileUploadResponse, error) {
	file, err := s.fileUploadRepo.GetFilesByID(fileID)
	if err != nil {
		return nil, e.NotFound("file not found")
	}

	response := model.ToFileUploadResponse(file, nil)
	return response, nil
}
