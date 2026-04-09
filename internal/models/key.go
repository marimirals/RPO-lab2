package models

import "time"

type Key struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    KeyValue    string    `gorm:"size:255;not null" json:"key_value"`
    Description string    `gorm:"size:255" json:"description"`
    CreatedAt   time.Time `json:"created_at"`
    Cards       []Card    `gorm:"foreignKey:KeyID" json:"cards,omitempty"`
}

type CreateKeyRequest struct {
    KeyValue    string `json:"key_value" binding:"required"`
    Description string `json:"description"`
}

type UpdateKeyRequest struct {
    KeyValue    string `json:"key_value"`
    Description string `json:"description"`
}