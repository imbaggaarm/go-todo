package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Title       string     `json:"title"`
	Time        time.Time  `json:"time"`
	CompletedAt *time.Time `json:"completed_at" gorm:"index"`
	IsDone      bool       `json:"is_done" gorm:"default:0"`
	UserID      uint       `json:"user_id" gorm:"unique;not null"`
	Note        string     `json:"note" gorm:"type:text"`
}
