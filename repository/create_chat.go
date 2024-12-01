package repository

import (
    "chat-app/config"
    "chat-app/models"
    "context"
    "github.com/jackc/pgx/v4"
    "fmt"
)

// Создать чат между двумя пользователями
// В функции GetChat
func GetChat(userID1, userID2 uint) (*models.Chat, error) {
    var chat models.Chat
    query := `SELECT id, user_id1, user_id2 FROM chats WHERE (user_id1 = $1 AND user_id2 = $2) OR (user_id1 = $2 AND user_id2 = $1)`

    fmt.Printf("Executing query: %s with params: %d, %d\n", query, userID1, userID2)
    
    row := config.DB.QueryRow(context.Background(), query, userID1, userID2)
    
    err := row.Scan(&chat.ID, &chat.UserID1, &chat.UserID2)
    if err != nil {
        fmt.Printf("Error executing query: %v\n", err)
        if err == pgx.ErrNoRows {
            return nil, nil // Чат не найден
        }
        return nil, err
    }
    
    fmt.Printf("Chat found: %+v\n", chat)
    return &chat, nil
}


// В функции CreateChat
func CreateChat(userID1, userID2 uint) (*models.Chat, error) {
    query := `INSERT INTO chats (user_id1, user_id2) VALUES ($1, $2) RETURNING id`
    var chatID uint
    row := config.DB.QueryRow(context.Background(), query, userID1, userID2)
    
    err := row.Scan(&chatID)
    if err != nil {
        fmt.Printf("Error creating chat: %v\n", err)
        return nil, err
    }

    chat := &models.Chat{
        ID: chatID,
        UserID1: userID1,
        UserID2: userID2,
    }
    fmt.Printf("Chat created: %+v\n", chat)
    return chat, nil
}
