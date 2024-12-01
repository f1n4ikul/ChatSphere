package controllers

import (
    "chat-app/repository"
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
    "chat-app/utils"
    "fmt"
)

// Начать чат с пользователем
// Начать чат с пользователем
func StartChat(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("userID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    currentUserID, err := utils.GetUserFromSession(c)
    if err != nil || currentUserID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    fmt.Printf("Creating chat for currentUserID: %d with userID: %d\n", currentUserID, userID)

    // Проверяем, существует ли уже чат с этим пользователем
    chat, err := repository.GetChat(uint(currentUserID), uint(userID))
    if err != nil {
        fmt.Printf("Error getting chat: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get chat"})
        return
    }

    if chat == nil {
        fmt.Println("No chat found, creating new chat...")
        chat, err = repository.CreateChat(uint(currentUserID), uint(userID))
        if err != nil {
            fmt.Printf("Error creating chat: %v\n", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create chat"})
            return
        }
    }

    if chat == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Chat is nil"})
        return
    }

    fmt.Printf("Chat created with ID: %d\n", chat.ID) // Логируем успешное создание чата
    c.JSON(http.StatusOK, gin.H{"chat_id": chat.ID})
}



// Отправить сообщение в чат
func SendMessage(c *gin.Context) {
    chatID, _ := strconv.Atoi(c.Param("chatID"))
    var input struct {
        Content string `json:"content"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // userID := 1 // ID текущего пользователя, который отправляет сообщение
    userID, err := utils.GetUserFromSession(c)
    if err != nil || userID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Отправляем сообщение
    _, err = repository.SendMessage(uint(chatID), uint(userID), input.Content)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
        return
    }

    c.Redirect(http.StatusFound, "/chats/"+strconv.Itoa(chatID))
}

// Получить все сообщения из чата
func GetMessages(c *gin.Context) {
    chatID, _ := strconv.Atoi(c.Param("chatID"))

    // Получаем все сообщения из чата
    messages, err := repository.GetMessages(uint(chatID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"messages": messages})
}