package models

import "time"

type Employee struct {
	ID        uint      `gorm:"primaryKey"`
	FirstName string    `gorm:"size:255;not null"`
	LastName  string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;unique;not null"`
	Phone     string    `gorm:"size:20"`
	Address   string    `gorm:"size:255"`
	Position  string    `gorm:"size:255"`
	Salary    float64   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}