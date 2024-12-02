package models

import "time"

// Message представляет сообщение в чате
type Message struct {
    ID             uint      // ID сообщения
    ChatID         uint      // ID чата
    SenderID       uint      // ID отправителя
    SenderUsername string    // Имя отправителя
    Content        string    // Содержимое сообщения
    Status         string    // Статус: "sent", "read"
    CreatedAt      time.Time // Время отправки
}