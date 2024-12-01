// package models

// import (
//     "github.com/jinzhu/gorm"
    
// )

// // Chat представляет чат между двумя пользователями
// type Chat struct {
//     gorm.Model
// 	ID uint   `gorm:"primary_key"` // ID чата
//     UserID1 uint   `gorm:"not null"` // ID первого пользователя
//     UserID2 uint   `gorm:"not null"` // ID второго пользователя
//     Messages []Message `gorm:"foreignKey:ChatID"`
// }

package models

// Chat представляет чат между двумя пользователями
type Chat struct {
    ID     uint   // ID чата
    UserID1 uint  // ID первого пользователя
    UserID2 uint  // ID второго пользователя
    Messages []Message 
}