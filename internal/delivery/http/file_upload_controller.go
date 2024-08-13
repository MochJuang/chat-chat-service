package http

import (
	e "chat-service/internal/exception"
	"chat-service/internal/model"
	"chat-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type FileUploadController struct {
	FileUploadService service.FileUploadService
}

func NewFileUploadController(service service.FileUploadService) *FileUploadController {
	return &FileUploadController{FileUploadService: service}
}

func (h *FileUploadController) UploadFile(c *fiber.Ctx) error {
	fileDTO := new(model.CreateFileUploadRequest)
	if err := c.BodyParser(fileDTO); err != nil {
		return e.Validation(err)
	}

	file, err := h.FileUploadService.UploadFile(fileDTO)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(file)
}

func (h *FileUploadController) GetFileByID(c *fiber.Ctx) error {
	fileID, err := c.ParamsInt("fileID")
	if err != nil {
		return e.Internal(err)
	}

	file, err := h.FileUploadService.GetFileByID(uint(fileID))
	if err != nil {
		return err
	}

	return c.JSON(file)
}
