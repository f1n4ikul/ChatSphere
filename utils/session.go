package utils

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret_key")) // Настройка секретного ключа для сессий

// Сохранение идентификатора пользователя в сессии
func SaveUserSession(c *gin.Context, userID uint) error {
    session, err := store.Get(c.Request, "user-session")
    if err != nil {
        return err
    }

    session.Values["userID"] = userID
    err = session.Save(c.Request, c.Writer)
    return err
}

// Получение идентификатора пользователя из сессии
func GetUserFromSession(c *gin.Context) (uint, error) {
    session, err := store.Get(c.Request, "user-session")
    if err != nil {
        return 0, err
    }

    userID, ok := session.Values["userID"].(uint)
    if !ok {
        return 0, nil // Нет сессии
    }

    return userID, nil
}

// Удаление сессии
func DestroySession(c *gin.Context) error {
    session, err := store.Get(c.Request, "user-session")
    if err != nil {
        return err
    }

    delete(session.Values, "userID")
    return session.Save(c.Request, c.Writer)
}