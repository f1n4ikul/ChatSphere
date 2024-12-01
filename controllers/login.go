package controllers

import (
    "chat-app/repository"
    "chat-app/utils"
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

// Обработка POST запроса для входа
func Login(c *gin.Context) {
    var input LoginInput

    // Получаем данные из запроса
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
        return
    }

    // Получаем пользователя по email
    user, err := repository.GetUserByEmail(input.Email)
    if err != nil || user == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
        return
    }

    // Проверка пароля
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
        return
    }

    // Сохраняем информацию о пользователе в сессии
    err = utils.SaveUserSession(c, user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
        return
    }

    // Редирект на Dashboard после успешной авторизации
   c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
    })
}