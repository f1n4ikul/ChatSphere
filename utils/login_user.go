package utils

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

// Функция для сохранения ID пользователя в сессии
func loginUser(c *gin.Context, userID uint) {
    session := sessions.Default(c)
    session.Set("user_id", userID)
    session.Save()
}