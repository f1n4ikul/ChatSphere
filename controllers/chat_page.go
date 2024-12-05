package controllers

import (
    "chat-app/utils"
    "chat-app/repository"
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
)

// Получить страницу чата с сообщениями
func ChatPage(c *gin.Context) {
    // Преобразуем chatID из параметра URL
    chatID, err := strconv.Atoi(c.Param("chatID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
        return
    }

    // Получаем userID из сессии
    userID, err := utils.GetUserFromSession(c)
    if err != nil || userID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    currentUser, err := repository.GetUserByID(uint(userID)) // Получаем имя текущего пользователя
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current user"})
        return
    }

    // Получаем все сообщения из чата
    messages, err := repository.GetMessages(uint(chatID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
        return
    }

    // Отправляем HTML-страницу с сообщениями
    c.HTML(http.StatusOK, "chat.html", gin.H{
        "ChatID":     chatID,
        "CurrentUser": currentUser.Username, // Здесь передаем только имя пользователя
        "Messages":   messages,
    })
}