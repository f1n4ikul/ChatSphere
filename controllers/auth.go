package controllers

import (   
    "net/http"
    "github.com/gin-gonic/gin"
    "chat-app/utils"
)

// Выход из системы
func Logout(c *gin.Context) {
    // Удаляем сессию
    err := utils.DestroySession(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out"})
        return
    }


    c.Redirect(http.StatusFound, "/login")
}

// Главная страница
func Index(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
}

// Страница входа
func LoginPage(c *gin.Context) {
    c.HTML(http.StatusOK, "login.html", nil)
}

// Страница регистрации
func RegisterPage(c *gin.Context) {
    c.HTML(http.StatusOK, "register.html", nil)
}
