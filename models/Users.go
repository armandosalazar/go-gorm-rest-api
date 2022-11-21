package models

import "time"

type User struct {
	ID       uint      `gorm:"type:integer;autoIncrement;primaryKey"`
	Email    string    `gorm:"not null;size:255;unique"`
	Password string    `gorm:"not null;size:255"`
	CreateAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Posts    []Post
}
