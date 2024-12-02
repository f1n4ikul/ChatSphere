package main

import (
    "chat-app/config"
    "chat-app/controllers"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/sessions"
    "log"
    "net/http"
)

// CookieStore для сессий
var store = sessions.NewCookieStore([]byte("your-secret-key"))

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Используем middleware для сессий с Gorilla sessions
    router.Use(func(c *gin.Context) {
        session, err := store.Get(c.Request, "mysession")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
            return
        }
        c.Set("session", session)
        c.Next()
    })

    // Рендеринг HTML-шаблонов
    router.LoadHTMLGlob("templates/*")

    // Главная страница
    router.GET("/", controllers.Index)

    // Страница для чата
    router.GET("/chats/:chatID", controllers.ChatPage)

    // Авторизация и регистрация
    router.GET("/login", controllers.LoginPage)
    router.POST("/login", controllers.Login)
    router.POST("/logout", controllers.Logout)

    router.GET("/register", controllers.RegisterPage)
    router.POST("/register", controllers.Register)

    // Дашборд
    router.GET("/dashboard", controllers.Dashboard)

    // Создание чата
    router.POST("/start-chat/:userID", controllers.StartChat)  // Новый маршрут для создания чата

    // Отправить сообщение
    router.POST("/chats/:chatID/messages", controllers.SendMessage)

    router.GET("/chat/:chatID", controllers.ChatPage)
    router.GET("/ws/chat/:chatID", controllers.WebSocketHandler)

    return router
}

func main() {
    config.InitDB()
    defer config.CloseDB()  // Закрываем подключение при завершении работы

    r := SetupRouter()
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
