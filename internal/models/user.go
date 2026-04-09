package models

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    Login        string    `gorm:"uniqueIndex;size:50;not null" json:"login"`
    Name         string    `gorm:"size:100;not null" json:"name"`
    PasswordHash string    `gorm:"size:255;not null" json:"-"` // не возвращаем пароль
    IsAdmin      bool      `gorm:"default:false" json:"is_admin"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

// CreateUserRequest - запрос на создание пользователя
type CreateUserRequest struct {
    Login    string `json:"login" binding:"required"`
    Name     string `json:"name" binding:"required"`
    Password string `json:"password" binding:"required,min=6"`
    IsAdmin  bool   `json:"is_admin"`
}

// UpdateUserRequest - запрос на обновление
type UpdateUserRequest struct {
    Name    string `json:"name"`
    IsAdmin *bool  `json:"is_admin"`
}

// LoginRequest - запрос на вход
type LoginRequest struct {
    Login    string `json:"login" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// LoginResponse - ответ при входе
type LoginResponse struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    User         User   `json:"user"`
}