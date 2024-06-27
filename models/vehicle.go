package models

import "time"

type Vehicle struct {
	ID          uint      `gorm:"primaryKey"`
	LicensePlate string   `gorm:"size:20;unique;not null"`
	Model       string    `gorm:"size:255;not null"`
	Brand       string    `gorm:"size:255;not null"`
	Year        int       `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}