package repository

import (
    "chat-app/models"
    "chat-app/config"
    "context"
)

// Получить всех пользователей, кроме текущего
func GetAllUsersExcept(userID uint) ([]models.User, error) {
    var users []models.User
    query := `SELECT id, username, email FROM users WHERE id != $1`
    rows, err := config.DB.Query(context.Background(), query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Username, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}
