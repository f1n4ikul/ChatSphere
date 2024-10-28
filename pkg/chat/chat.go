package chat

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v4"
    "net/http"
    "strconv"
    "time"
)

type Message struct {
    ID        int       `json:"id"`
    ChatID    int       `json:"chat_id"`
    SenderID  int       `json:"sender_id"`
    Text      string    `json:"text"`
    CreatedAt time.Time `json:"created_at"`
}

// Определите структуры для чата и сообщения
type CreateChatRequest struct {
    User1ID int `json:"user1_id"`
    User2ID int `json:"user2_id"`
}

// Определите структуры для чата и сообщения
type Chat struct {
    ID        int       `json:"id"`
    User1ID   int       `json:"user1_id"`
    User2ID   int       `json:"user2_id"`
    CreatedAt time.Time `json:"created_at"`
}

// Подключение к базе данных
func connectDB() (*pgx.Conn, error) {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost/Analogue_discord")
    return conn, err
}

// Создание нового чата
func CreateChat(c *gin.Context) {
    var req CreateChatRequest

    // Прочитайте данные из JSON
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
        return
    }

    user1ID := req.User1ID
    user2ID := req.User2ID

    // Подключитесь к базе данных
    conn, err := connectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось подключиться к базе данных"})
        return
    }
    defer conn.Close(context.Background())

    // Сохраните чат в базе данных
    var chatID int
    err = conn.QueryRow(context.Background(), "INSERT INTO chats (user1_id, user2_id) VALUES ($1, $2) RETURNING id", user1ID, user2ID).Scan(&chatID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать чат"})
        return
    }

    chat := Chat{
        ID:        chatID,
        User1ID:   user1ID,
        User2ID:   user2ID,
        CreatedAt: time.Now(),
    }

    c.JSON(http.StatusCreated, chat)
}

// Получение сообщений для чата
func GetMessages(c *gin.Context) {
    chatIDStr := c.Param("userID") // Измените на chatId, если используете его в URL

    // Преобразование строкового идентификатора чата в целое число
    chatID, err := strconv.Atoi(chatIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор чата"})
        return
    }

    // Подключитесь к базе данных
    conn, err := connectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось подключиться к базе данных"})
        return
    }
    defer conn.Close(context.Background())

    // Получите сообщения из базы данных
    rows, err := conn.Query(context.Background(), "SELECT id, chat_id, sender_id, text, created_at FROM messages WHERE chat_id = $1", chatID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить сообщения"})
        return
    }
    defer rows.Close()

    var messages []Message
    for rows.Next() {
        var message Message
        if err := rows.Scan(&message.ID, &message.ChatID, &message.SenderID, &message.Text, &message.CreatedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при чтении данных сообщения"})
            return
        }
        messages = append(messages, message)
    }

    c.JSON(http.StatusOK, messages)
}

// Отправка нового сообщения
func SendMessage(c *gin.Context) {
    chatIDStr := c.Param("chatID")
    senderIDStr := c.PostForm("sender_id")
    text := c.PostForm("text")

    // Преобразование строковых идентификаторов в целые числа
    chatID, err := strconv.Atoi(chatIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор чата"})
        return
    }
    senderID, err := strconv.Atoi(senderIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор отправителя"})
        return
    }

    // Подключитесь к базе данных
    conn, err := connectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось подключиться к базе данных"})
        return
    }
    defer conn.Close(context.Background())

    // Сохраните сообщение в базе данных
    var messageID int
    err = conn.QueryRow(context.Background(), "INSERT INTO messages (chat_id, sender_id, text) VALUES ($1, $2, $3) RETURNING id", chatID, senderID, text).Scan(&messageID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось отправить сообщение"})
        return
    }

    message := Message{
        ID:        messageID,
        ChatID:    chatID,
        SenderID:  senderID,
        Text:      text,
        CreatedAt: time.Now(),
    }

    c.JSON(http.StatusCreated, message)
}