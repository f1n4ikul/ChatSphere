package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "analogue_discord/pkg/auth" 
    "analogue_discord/pkg/chat" 
)

func main() {
    router := gin.Default()

    // Включение CORS
    router.Use(cors.Default())

    // Обслуживание статических файлов
    router.Static("/static", "./web/static")

    // Настройка маршрутов API
    router.POST("/api/register", auth.Register)
    router.POST("/api/login", auth.Login)
    router.GET("/api/user/:username", auth.Account)
    
    // Обслуживание чатов
    router.POST("/api/chats", chat.CreateChat) // Создание нового чата
    router.GET("/api/chats/:chatID", chat.GetMessages) // Получение сообщений для чата
    router.GET("/api/chats/:chatID/messages", chat.GetMessages) // Получение сообщений для чата
    router.POST("/api/chats/:chatID/messages", chat.SendMessage) // Отправка нового сообщения

    router.NoRoute(func(c *gin.Context) {
        c.File("./web/public/index.html")
    })

    router.Run(":8080")
}
