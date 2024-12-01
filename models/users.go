package models

// import "github.com/jinzhu/gorm"

type User struct {
    // gorm.Model
    ID uint `json:"id"`
    Email    string `json:"email"`
    Username string `json:"username"`
    Name     string `json:"name"`
    Surname  string `json:"surname"`
    Password string `json:"password"`
}
