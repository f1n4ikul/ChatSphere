package controllers

import (
    "chat-app/repository"
    "chat-app/utils"
    "net/http"
    "github.com/gin-gonic/gin"
)

// Получение информации о текущем пользователе
func Dashboard(c *gin.Context) {
    // Получаем ID пользователя из сессии
    userID, err := utils.GetUserFromSession(c)
    if err != nil || userID == 0 {
        // c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Redirect(http.StatusFound, "/login")
        return
    }

    // Получаем данные пользователя из базы
    user, err := repository.GetUserByID(userID)
    if err != nil || user == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user data"})
        return
    }

	users, err := repository.GetAllUsersExcept(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
        return
    }

    // Отправляем данные пользователя на фронт
    c.HTML(http.StatusOK, "dashboard.html", gin.H{
        "name":    user.Name,
        "email":   user.Email,
        "username": user.Username,
		"Users":   users,

    })
	
}