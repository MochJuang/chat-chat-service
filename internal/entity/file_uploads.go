package entity

import (
	"time"
)

type FileUpload struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Type      string    `gorm:"not null"` // Type can be something like "profile_picture", "attachment", etc.
	RefID     uint      `gorm:"not null"` // Reference ID for linking with other tables (e.g., message, conversation)
	FileURL   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

const (
	FileUploadTypeProfilePicture    = "PROFILE_PICTURE"
	FileUploadTypeMessageAttachment = "MESSAGE_ATTACHMENT"
)
