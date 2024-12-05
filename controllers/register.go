package controllers

import (
	"chat-app/models"
	"chat-app/repository"
	"chat-app/utils"
	"fmt"
	"log"
	"net/http"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/gin-gonic/gin"
)

const recaptchaSecretKey string = "6Lcc9ZIqAAAAAGGyWl3_g02b5yPYHdFERbXSFx6J" // Ваш секретный ключ

func init() {
	// Инициализация reCAPTCHA с секретным ключом
	recaptcha.Init(recaptchaSecretKey)
	log.Println("reCAPTCHA initialized successfully.")
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка капчи
	if !verifyCaptcha(user.Captcha, c.ClientIP()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Captcha verification failed"})
		return
	}

	// Хэшируем пароль
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword

	// Создание пользователя в базе данных
	if err := repository.CreateUser(&user); err != nil {
		fmt.Errorf("Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Функция для проверки капчи
func verifyCaptcha(captchaResponse string, clientIP string) bool {
	// Используем recaptcha.Confirm для проверки ответа
	success, err := recaptcha.Confirm(clientIP, captchaResponse)
	if err != nil {
		log.Println("Error verifying captcha:", err)
		return false
	}

	if !success {
		log.Println("Captcha verification failed.")
		return false
	}

	return true
}
