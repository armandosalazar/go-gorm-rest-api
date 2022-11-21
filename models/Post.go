package models

import (
	"time"
)

type Post struct {
	Content   string    `gorm:"not null;size:255"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UserID    uint      `gorm:"type:integer;foreignKey:ID"`
}
