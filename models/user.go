package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"not null;unique;" json:"email"`
	Password  string    `gorm:"not null;" json:"password"`
	Role      string    `gorm:"default:'Public'" json:"role"`
	Status    string    `gorm:"default:'Active'" json:"status"`
	Company   string    `gorm:"default:ITD" json:"company"`
	CreatedAt time.Time ``
	UpdatedAt time.Time ``
}
