package auth

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v4"
    "net/http"
)

// RegisterRequest - структура для получения данных регистрации
type RegisterRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Подключение к базе данных
func connectDB() (*pgx.Conn, error) {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost/Analogue_discord") // Проверьте строку подключения
    return conn, err
}

// Register - обработчик для регистрации
func Register(c *gin.Context) {
    var req RegisterRequest

    // Декодируем JSON из тела запроса
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Подключитесь к базе данных
    conn, err := connectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось подключиться к базе данных: " + err.Error()})
        return
    }
    defer conn.Close(context.Background())

    // Проверяем, существует ли уже пользователь
    var exists bool
    err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", req.Username).Scan(&exists)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке пользователя: " + err.Error()})
        return
    }

    if exists {
        c.JSON(http.StatusConflict, gin.H{"error": "Пользователь с таким именем уже существует"})
        return
    }

    // Сохраняем пользователя (не храните пароли в открытом виде!)
    _, err = conn.Exec(context.Background(), "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", req.Username, req.Email, req.Password) // Используйте хэширование пароля
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пользователя: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

// Login - обработчик для входа
func Login(c *gin.Context) {
    var req RegisterRequest // Можно использовать ту же структуру, если поля совпадают

    // Декодируем JSON из тела запроса
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Подключитесь к базе данных
    conn, err := connectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось подключиться к базе данных: " + err.Error()})
        return
    }
    defer conn.Close(context.Background())

    // Проверяем учетные данные пользователя
    var storedPassword string
    err = conn.QueryRow(context.Background(), "SELECT password FROM users WHERE username=$1", req.Username).Scan(&storedPassword)
    if err != nil {
        if err == pgx.ErrNoRows {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке учетных данных: " + err.Error()})
        }
        return
    }

    // Здесь также следует использовать хэширование пароля для сравнения
    if storedPassword != req.Password { // Используйте хэширование пароля вместо простого сравнения
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}

// Account - обработчик для получения информации о пользователе
func Account(c *gin.Context) {
    username := c.Param("username")

    // Подключитесь к базе данных
    conn, err := connectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось подключиться к базе данных: " + err.Error()})
        return
    }
    defer conn.Close(context.Background())

    // Проверяем, существует ли пользователь
    var exists bool
    err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке пользователя: " + err.Error()})
        return
    }

    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"username": username})
}