package model

type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"` // В реальном приложении не храните пароль в открытом виде!
    Email    string `json:"email"`
    
    
}