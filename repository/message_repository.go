package repository

import (
    "chat-app/config"
    "chat-app/models"
    "context"
    "time"
    "fmt"
    "log"
)

// Отправить сообщение в чат
func SendMessage(chatID uint, senderID uint, content string) (*models.Message, error) {
    query := `INSERT INTO messages (chat_id, sender_id, content) VALUES ($1, $2, $3) RETURNING id, created_at`
    row := config.DB.QueryRow(context.Background(), query, chatID, senderID, content)

    if chatID == 0 {
        return nil, fmt.Errorf("chatID cannot be 0")
    }

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
    query := `
        SELECT m.id, m.chat_id, m.sender_id, m.content, m.created_at, m.status, u.username 
        FROM messages m
        JOIN users u ON m.sender_id = u.id
        WHERE m.chat_id = $1 
        ORDER BY m.created_at ASC`
    
    rows, err := config.DB.Query(context.Background(), query, chatID)
    if err != nil {
        log.Printf("Error executing query: %v", err)
        return nil, err
    }

    if chatID == 0 {
        return nil, fmt.Errorf("chatID cannot be 0")
    }

    defer rows.Close()

    for rows.Next() {
        var message models.Message
        err := rows.Scan(&message.ID, &message.ChatID, &message.SenderID, &message.Content, &message.CreatedAt, &message.Status, &message.SenderUsername)
        if err != nil {
            log.Printf("Error scanning row: %v", err)
            return nil, err
        }
        messages = append(messages, message)
    }

    if err := rows.Err(); err != nil {
        log.Printf("Error scanning rows: %v", err)
        return nil, err
    }

    return messages, nil
}


// Сохранить новое сообщение в базе данных
func SaveMessage(chatID, senderID uint, content string) (*models.Message, error) {
    query := `INSERT INTO messages (chat_id, sender_id, content, status) 
              VALUES ($1, $2, $3, 'sent') RETURNING id, created_at`
    var messageID uint
    var createdAt time.Time

    row := config.DB.QueryRow(context.Background(), query, chatID, senderID, content)
    err := row.Scan(&messageID, &createdAt)
    if err != nil {
        return nil, err
    }

    if chatID == 0 {
        return nil, fmt.Errorf("chatID cannot be 0")
    }

    message := &models.Message{
        ID:        messageID,
        ChatID:    chatID,
        SenderID:  senderID,
        Content:   content,
        Status:    "unread",
        CreatedAt: createdAt,
    }

    return message, nil
}
