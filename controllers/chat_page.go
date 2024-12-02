package controllers

import (
    // "chat-app/models"
    "chat-app/utils"
    "chat-app/repository"
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
)

// Получить страницу чата с сообщениями
func ChatPage(c *gin.Context) {
    chatID, err := strconv.Atoi(c.Param("chatID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
        return
    }

    userID, err := utils.GetUserFromSession(c)
    if err != nil || userID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Получаем все сообщения из чата
    messages, err := repository.GetMessages(uint(chatID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
        return
    }

    c.HTML(http.StatusOK, "chat.html", gin.H{
        "ChatID":  chatID,
        "Messages": messages,
    })
}
