package domain

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255)" form:"name" json:"name"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex" form:"email" json:"email"`
	Password  string    `gorm:"type:varchar(255)" form:"password" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"index" form:"user_id" json:"user_id"`
	Title       string    `gorm:"type:varchar(255)" form:"title" json:"title"`
	Description string    `gorm:"type:text" form:"description" json:"description"`
	Status      string    `gorm:"type:varchar(50);default:'pending'" form:"status" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Login struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ResponseModels struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
