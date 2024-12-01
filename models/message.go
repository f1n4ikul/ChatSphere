package models

import "gorm.io/gorm"

// Message представляет сообщение в чате
type Message struct {
    gorm.Model
    ChatID   uint   `gorm:"not null"`
    SenderID uint   `gorm:"not null"`
    Content  string `gorm:"not null"`
}