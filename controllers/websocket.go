package controllers

import (
    "chat-app/repository"
    "chat-app/models"
    "net/http"
    "github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
    "strconv"
    "log"
)

// Создаем глобальную переменную для хранения подключенных клиентов
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

// WebSocket обработчик
func WebSocketHandler(c *gin.Context) {
    chatID, err := strconv.Atoi(c.Param("chatID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
        return
    }

    // Устанавливаем WebSocket соединение
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    clients[conn] = true

    // Отправляем последние сообщения при подключении
    messages, err := repository.GetMessages(uint(chatID))
    if err != nil {
        log.Println("Failed to get messages:", err)
        return
    }
    
    for _, message := range messages {
        err := conn.WriteJSON(message)
        if err != nil {
            log.Println("Error sending message:", err)
            return
        }
    }

    // Чтение сообщений от клиента и их отправка всем подключенным
    for {
        var message models.Message
        err := conn.ReadJSON(&message)
        if err != nil {
            delete(clients, conn)
            break
        }

        // Отправляем сообщение всем клиентам
        broadcast <- message
    }
}

// Отправка сообщений всем подключенным клиентам
func handleMessages() {
    for {
        message := <-broadcast
        for client := range clients {
            err := client.WriteJSON(message)
            if err != nil {
                log.Println("Error sending message:", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func init() {
    go handleMessages()
}
