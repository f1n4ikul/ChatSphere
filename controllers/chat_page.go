package controllers

import (
    "chat-app/models"
    "chat-app/repository"
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
)

// Страница чата
func ChatPage(c *gin.Context) {
    chatID, _ := strconv.Atoi(c.Param("chatID"))
    messages, err := repository.GetMessages(uint(chatID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
        return
    }

    // Создаем структуру с данными для шаблона
    data := struct {
        ChatID   uint
        Messages []models.Message
    }{
        ChatID:   uint(chatID),
        Messages: messages,
    }

    c.HTML(http.StatusOK, "chat.html", data)
}
