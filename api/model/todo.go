package model

import (
	"errors"
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

func GetUserTodos(userID, offSet uint64) (*[]Todo, error) {
	var results *[]Todo
	err := GetDB().
		Where("user_id = ?", userID).
		Order("id").
		Offset(offSet).
		Limit(20).
		Find(results).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("connection error. Retry later")
	}
	return results, nil
}
