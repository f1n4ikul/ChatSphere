package model

type Message struct {
    ID      uint   `json:"id"`
    UserID  uint   `json:"user_id"`
    Content string `json:"content"`
}