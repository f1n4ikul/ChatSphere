package config

import (
    "fmt"
    "log"
    "github.com/jackc/pgx/v4"
    "context"
)

var DB *pgx.Conn

// InitDB инициализирует подключение к базе данных
func InitDB() {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost/create_app_chat")
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }

    DB = conn
    fmt.Println("Connected to PostgreSQL")
}

// Закрытие подключения
func CloseDB() {
    if DB != nil {
        DB.Close(context.Background())
    }
}
