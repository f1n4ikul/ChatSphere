package auth

import (
	"net/url"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
)

func TestRegister(t *testing.T) {
    router := gin.Default()
    router.POST("/api/register", Register)
	
    // Создаем запрос
    req, _ := http.NewRequest("POST", "/api/register", nil)
    req.PostForm = url.Values{
        "username": {"testuser"},
        "password": {"password123"},
        "email":    {"test@example.com"},
    }
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Проверяем статус ответа
    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}