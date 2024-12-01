// package routes

// import (
//     "chat-app/controllers"
//     "github.com/gin-gonic/gin"
// )

// func SetupRouter() *gin.Engine {
//     router := gin.Default()

//     // Настройка статических файлов
//     router.Static("/static", "./static")

//     // Рендеринг HTML-шаблонов
//     router.LoadHTMLGlob("templates/*")

//     // Главная страница
//     router.GET("/", controllers.Index)

//     // Страница для чата
//     router.GET("/chats/:chatID", controllers.ChatPage)

//     // Авторизация и регистрация
//     router.GET("/login", controllers.LoginPage)
//     router.POST("/login", controllers.Login)

//     router.GET("/register", controllers.RegisterPage)
//     router.POST("/register", controllers.Register)

//     // Отправить сообщение
//     router.POST("/chats/:chatID/messages", controllers.SendMessage)

//     return router
// }


