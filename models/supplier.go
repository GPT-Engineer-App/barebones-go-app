package models

import "time"

type Supplier struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;unique;not null"`
	Phone     string    `gorm:"size:20"`
	Address   string    `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}