package controllers

import (   
    "net/http"
    "github.com/gin-gonic/gin"
    
)


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
