package models

import "time"

type Trip struct {
	ID        uint      `gorm:"primaryKey"`
	RouteID   uint      `gorm:"not null"`
	VehicleID uint      `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}