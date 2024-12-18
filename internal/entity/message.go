package entity

import (
	"time"
)

type Message struct {
	ID             uint         `gorm:"primaryKey"`
	UUID           string       `gorm:"unique;not null"`
	SenderID       uint         `gorm:"not null"`
	Sender         User         `gorm:"foreignKey:SenderID"`
	ConversationID uint         `gorm:"not null"`
	Conversation   Conversation `gorm:"foreignKey:ConversationID"`
	Content        string       `gorm:"type:text;not null"`
	SendAt         time.Time    `gorm:"autoCreateTime"`
}
