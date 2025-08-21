package models

import "time"

type User struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string     `gorm:"type:varchar(100);not null" json:"name"`
	Email        string     `gorm:"type:varchar(191);uniqueIndex;not null" json:"email"`
	Password     string     `gorm:"type:varchar(191);not null" json:"-"`
	Role         string     `gorm:"type:varchar(20);default:'user'" json:"role"`
	TokenVersion int        `gorm:"default:1" json:"-"` // <— INI YANG KURANG
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index" json:"-"`
}
