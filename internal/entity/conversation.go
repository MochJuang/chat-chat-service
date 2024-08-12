package entity

import (
	"time"
)

type Conversation struct {
	ID           uint      `gorm:"primaryKey"`
	Participants []User    `gorm:"many2many:conversation_participants"`
	Messages     []Message `gorm:"foreignKey:ConversationID"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
