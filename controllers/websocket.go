package controllers

import (
	"chat-app/repository"
	"chat-app/models"
	"context"
	"chat-app/utils"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
	"encoding/json"
	"chat-app/config"
	"log"
	"time"
)

// Глобальная переменная для хранения подключенных клиентов
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketManager struct {
	clients map[uint][]*websocket.Conn // Хранит клиентов, подключенных к каждому чату
	mutex   sync.Mutex
}

var manager = WebSocketManager{
	clients: make(map[uint][]*websocket.Conn),
}

// WebSocket обработчик
func WebSocketHandler(c *gin.Context) {
	// Получаем ID чата из параметров
	chatID, err := strconv.Atoi(c.Param("chatID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
		return
	}

	// Получаем userID из сессии
	userID, err := utils.GetUserFromSession(c)
	if err != nil || userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Устанавливаем WebSocket соединение
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	// Добавляем клиента в список активных WebSocket-соединений
	manager.mutex.Lock()
	manager.clients[uint(chatID)] = append(manager.clients[uint(chatID)], conn)
	manager.mutex.Unlock()

	// Ожидаем получения сообщений от клиента
	for {
		// Чтение сообщения от клиента
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("Unexpected WebSocket closure:", err)
			}
			break
		}

		var msg struct {
			Content string `json:"content"`
			ChatID  string `json:"chat_id"`
		}

		// Десериализация сообщения
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}

		// Преобразуем chatID
		parsedChatID, err := strconv.ParseUint(msg.ChatID, 10, 32)
		if err != nil {
			log.Println("Error converting chat_id to uint:", err)
			continue
		}

		// Сохраняем сообщение в базе данных
		savedMessage, err := repository.SaveMessage(uint(parsedChatID), uint(userID), msg.Content)
		if err != nil {
			log.Println("Error saving message:", err)
			continue
		}

		// Получаем пользователя для имени отправителя
		user, err := repository.GetUserByID(uint(userID))
		if err != nil {
			log.Println("Error fetching user:", err)
			continue
		}

		// Создаем отформатированное сообщение
		formattedMessage := models.Message{
			ID:             savedMessage.ID,
			SenderUsername: user.Username,
			Content:        savedMessage.Content,
			CreatedAt:      savedMessage.CreatedAt,
			Status:        "unread",
		}

		// Сериализуем сообщение в JSON
		messageBytes, err := json.Marshal(formattedMessage)
		if err != nil {
			log.Println("Error marshaling message:", err)
			continue
		}

		// Обновляем статус сообщения на "прочитано"
		err = markMessageAsRead(savedMessage.ID)
		if err != nil {
			log.Println("Error marking message as read:", err)
		}

		// Отправляем сообщение всем подключённым пользователям
		manager.mutex.Lock()
		for _, client := range manager.clients[uint(parsedChatID)] {
			err := sendToClient(client, messageBytes)
			if err != nil {
				log.Println("Error sending message to client:", err)
				client.Close()
				// Удаляем клиента из списка после ошибки
				removeClientFromChat(client, uint(parsedChatID))
			}
		}
		manager.mutex.Unlock()

		// Отправляем подтверждение клиенту
		// conn.WriteJSON(savedMessage)
	}
}

// Функция для отправки сообщения клиенту с проверкой состояния соединения
func sendToClient(client *websocket.Conn, message []byte) error {
	client.SetWriteDeadline(time.Now().Add(time.Second * 10)) // Устанавливаем дедлайн на 10 секунд
	return client.WriteMessage(websocket.TextMessage, message)
}

// Удаление клиента из чата
func removeClientFromChat(client *websocket.Conn, chatID uint) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	for i, c := range manager.clients[chatID] {
		if c == client {
			manager.clients[chatID] = append(manager.clients[chatID][:i], manager.clients[chatID][i+1:]...)
			break
		}
	}
}

// Маркировка сообщения как "прочитанное"
func markMessageAsRead(messageID uint) error {
	query := `UPDATE messages SET status = 'read' WHERE id = $1`
	_, err := config.DB.Exec(context.Background(), query, messageID)
	return err
}

// Отправка сообщений всем подключенным клиентам
func handleMessages() {
	for {
		message := <-broadcast
		// Сериализуем сообщение в байтовый срез
		messageBytes, err := json.Marshal(message)
		if err != nil {
			log.Println("Error marshaling broadcast message:", err)
			continue
		}

		// Отправляем сообщение всем подключенным клиентам
		for client := range clients {
			err := sendToClient(client, messageBytes)
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
