package utils

import (
    "golang.org/x/crypto/bcrypt"
    // "github.com/gin-gonic/gin"
)

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashedPassword), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// func DestroyDeliteSession(c *gin.Context) error {
//     // Пример удаления cookie-сессии
//     c.SetCookie("session_id", "", -1, "/", "localhost", false, true)
//     return nil
// }