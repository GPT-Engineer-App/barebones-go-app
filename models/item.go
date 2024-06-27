package models

import "time"

type Item struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null"`
	Description string    `gorm:"size:255"`
	Price       float64   `gorm:"not null"`
	SupplierID  uint      `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}