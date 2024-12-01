package repository

import (
    "chat-app/models"
    "chat-app/config"
    "context"
    "github.com/jackc/pgx/v4"
    "errors"
)

// Создать нового пользователя
func CreateUser(user *models.User) error {
    query := `INSERT INTO users (email, username, name, surname, password_hash) VALUES ($1, $2, $3, $4, $5)`
    _, err := config.DB.Exec(context.Background(), query, user.Email, user.Username, user.Name, user.Surname, user.Password)
    return err
}

// Получить пользователя по email
func GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    query := `SELECT id, email, username, name, surname, password_hash FROM users WHERE email=$1`
    row := config.DB.QueryRow(context.Background(), query, email)

    err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Name, &user.Surname, &user.Password)
    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }

    return &user, nil
}


// Получить пользователя по ID
func GetUserByID(userID uint) (*models.User, error) {
    var user models.User
    // Выполняем SQL-запрос для получения данных пользователя по его ID
    query := `SELECT id, email, username, password_hash, name, surname FROM users WHERE id=$1`
    row := config.DB.QueryRow(context.Background(), query, userID)

    // Заполняем структуру данных пользователя
    err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Name, &user.Surname)
    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("user not found")
        }
        return nil, err
    }

    // Возвращаем данные пользователя
    return &user, nil
}