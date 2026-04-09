package models

import "time"

type Transaction struct {
    ID             uint      `gorm:"primaryKey" json:"id"`
    CardID         uint      `gorm:"not null;index" json:"card_id"`
    TerminalID     uint      `gorm:"not null;index" json:"terminal_id"`
    Amount         int64     `gorm:"not null" json:"amount"` // в копейках
    TransactionType string    `gorm:"size:20;default:'payment'" json:"transaction_type"`
    Status         string    `gorm:"size:20;default:'pending'" json:"status"`
    TransactionTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"transaction_time"`
    Card           *Card     `gorm:"foreignKey:CardID" json:"card,omitempty"`
    Terminal       *Terminal `gorm:"foreignKey:TerminalID" json:"terminal,omitempty"`
}

type CreateTransactionRequest struct {
    CardID     uint   `json:"card_id" binding:"required"`
    TerminalID uint   `json:"terminal_id" binding:"required"`
    Amount     int64  `json:"amount" binding:"required,gt=0"`
    Type       string `json:"type"`
}

// TerminalAuthRequest - запрос от терминала на авторизацию
type TerminalAuthRequest struct {
    CardNumber string `json:"card_number" binding:"required"`
    Amount     int64  `json:"amount" binding:"required,gt=0"`
    TerminalID uint   `json:"terminal_id" binding:"required"`
}

// TerminalAuthResponse - ответ терминалу
type TerminalAuthResponse struct {
    Authorized bool   `json:"authorized"`
    Message    string `json:"message"`
    Balance    int64  `json:"balance,omitempty"`
}