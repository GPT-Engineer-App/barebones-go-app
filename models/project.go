package models

import "time"

type Project struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null"`
	Description string    `gorm:"size:255"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     time.Time
	ClientID    uint      `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}