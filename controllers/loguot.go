package controllers

import (
    
    "github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
    "net/http"
)

// Обработчик для выхода из системы
func Logout(c *gin.Context) {
    // Получаем сессию из контекста
    session := c.MustGet("session").(*sessions.Session)

    // Удаляем информацию о пользователе из сессии
    session.Options.MaxAge = -1 // Устанавливаем время жизни сессии в минус, чтобы удалить cookie
    err := session.Save(c.Request, c.Writer) // Сохраняем изменения в сессии
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
        return
    }

    // Отправляем успешный ответ
    c.JSON(http.StatusOK, gin.H{
        "message": "Logout successful",
    })
}