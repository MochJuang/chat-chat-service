package entity

import (
	"time"
)

type Notification struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      string    `gorm:"unique;not null"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Message   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
