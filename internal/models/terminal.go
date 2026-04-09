package models

import "time"

type Terminal struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    SerialNumber string    `gorm:"uniqueIndex;size:50;not null" json:"serial_number"`
    Name         string    `gorm:"size:100;not null" json:"name"`
    Address      string    `gorm:"size:255" json:"address"`
    Location     string    `gorm:"size:255" json:"location"`
    IsActive     bool      `gorm:"default:true" json:"is_active"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type CreateTerminalRequest struct {
    SerialNumber string `json:"serial_number" binding:"required"`
    Name         string `json:"name" binding:"required"`
    Address      string `json:"address"`
    Location     string `json:"location"`
}

type UpdateTerminalRequest struct {
    Name     string `json:"name"`
    Address  string `json:"address"`
    Location string `json:"location"`
    IsActive *bool  `json:"is_active"`
}