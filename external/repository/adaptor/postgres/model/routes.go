package gorm_model

import "time"

type Route struct {
	ID          string `gorm:"primarykey"`
	Name        string `gorm:"not null"`
	OwnerId     string `gorm:"not null"`
	Description string
	Flow        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}