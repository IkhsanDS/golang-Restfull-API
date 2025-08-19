package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"` // ✅ char(36), bukan uuid
	Title     string         `gorm:"size:255;not null" json:"title"`
	Completed bool           `gorm:"type:tinyint(1);not null;default:0" json:"completed"` // ✅ tinyint(1) + default 0
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

//	func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
//		if t.ID == "" {
//			t.ID = uuid.New().String()
//		}
//		return nil
//	}
func (t *Todo) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}
