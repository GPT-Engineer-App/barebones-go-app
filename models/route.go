package models

import "time"

type Route struct {
	ID          uint      `gorm:"primaryKey"`
	Origin      string    `gorm:"size:255;not null"`
	Destination string    `gorm:"size:255;not null"`
	Distance    float64   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}