package repository

import (
    "chat-app/config"
    "chat-app/models"
    "context"
)

// Отправить сообщение в чат
func SendMessage(chatID uint, senderID uint, content string) (*models.Message, error) {
    query := `INSERT INTO messages (chat_id, sender_id, content) VALUES ($1, $2, $3) RETURNING id, created_at`
    row := config.DB.QueryRow(context.Background(), query, chatID, senderID, content)

    var message models.Message
    err := row.Scan(&message.ID, &message.CreatedAt)
    if err != nil {
        return nil, err
    }

    message.ChatID = chatID
    message.SenderID = senderID
    message.Content = content

    return &message, nil
}

// Получить все сообщения из чата
func GetMessages(chatID uint) ([]models.Message, error) {
    var messages []models.Message
    query := `SELECT id, chat_id, sender_id, content, created_at FROM messages WHERE chat_id = $1 ORDER BY created_at ASC`
    rows, err := config.DB.Query(context.Background(), query, chatID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var message models.Message
        err := rows.Scan(&message.ID, &message.ChatID, &message.SenderID, &message.Content, &message.CreatedAt)
        if err != nil {
            return nil, err
        }
        messages = append(messages, message)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return messages, nil
}
