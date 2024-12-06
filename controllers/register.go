package controllers

import (
	"chat-app/models"
	"chat-app/repository"
	"chat-app/utils"
	"fmt"
	"log"
	"bytes"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
)



const recaptchaSecretKey string = "6Lcc9ZIqAAAAAGGyWl3_g02b5yPYHdFERbXSFx6J" // Ваш секретный ключ

// Функция для проверки капчи
func verifyCaptcha(captchaResponse string, clientIP string) bool {
	url := "https://www.google.com/recaptcha/api/siteverify"
	data := fmt.Sprintf("secret=%s&response=%s&remoteip=%s", recaptchaSecretKey, captchaResponse, clientIP)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Println("Error creating POST request:", err)
		return false
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return false
	}
	defer resp.Body.Close()

	var result struct {
		Success    bool     `json:"success"`
		ErrorCodes []string `json:"error-codes"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error decoding CAPTCHA verification response:", err)
		return false
	}

	// Проверка на успешность капчи
	if result.Success {
		return true
	}
	log.Println("reCAPTCHA failed:", result.ErrorCodes)
	return false
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
		log.Printf("Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

