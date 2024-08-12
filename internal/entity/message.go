package entity

import (
	"time"
)

type Message struct {
	ID             uint         `gorm:"primaryKey"`
	SenderID       uint         `gorm:"not null"`
	Sender         User         `gorm:"foreignKey:SenderID"`
	ConversationID uint         `gorm:"not null"`
	Conversation   Conversation `gorm:"foreignKey:ConversationID"`
	Content        string       `gorm:"type:text;not null"`
	SendAt         time.Time    `gorm:"autoCreateTime"`
}
