package models

import "time"

type Payment struct {
	ID        uint      `gorm:"primaryKey"`
	Amount    float64   `gorm:"not null"`
	Date      time.Time `gorm:"not null"`
	ClientID  uint      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}