package models

import "time"

type Card struct {
    ID         uint       `gorm:"primaryKey" json:"id"`
    CardNumber string     `gorm:"uniqueIndex;size:20;not null" json:"card_number"`
    Balance    int64      `gorm:"default:0" json:"balance"` // храним в копейках
    IsBlocked  bool       `gorm:"default:false" json:"is_blocked"`
    OwnerName  string     `gorm:"size:100" json:"owner_name"`
    KeyID      *uint      `gorm:"index" json:"key_id"`
    Key        *Key       `gorm:"foreignKey:KeyID" json:"key,omitempty"`
    CreatedAt  time.Time  `json:"created_at"`
    UpdatedAt  time.Time  `json:"updated_at"`
}

type CreateCardRequest struct {
    CardNumber string `json:"card_number" binding:"required"`
    Balance    *int64 `json:"balance"`
    OwnerName  string `json:"owner_name"`
    KeyID      *uint  `json:"key_id"`
}

type UpdateCardRequest struct {
    Balance   *int64 `json:"balance"`
    IsBlocked *bool  `json:"is_blocked"`
    OwnerName string `json:"owner_name"`
    KeyID     *uint  `json:"key_id"`
}