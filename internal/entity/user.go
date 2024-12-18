package entity

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      string    `gorm:"unique;not null"`
	Name      string    `gorm:"not null"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
