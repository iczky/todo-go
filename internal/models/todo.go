package models

import "time"

type Todo struct {
	ID          uint      `gorm:"primary_key"`
	Title       string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	Priority    int       `gorm:"default:0"`
	Completed   bool      `gorm:"default:false"`
	DueDate     time.Time `gorm:""`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
