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

func (todo *Todo) CreateTodo() error {
	GetDB().Create(todo)
	if todo.ID <= 0 {
		return errors.New("create todo failed. Please retry")
	}
	return nil
}

func GetUserTodos(userID, offSet uint64) (*[]Todo, error) {
	var results []Todo
	err := GetDB().
		Where("user_id = ?", userID).
		Order("id").
		Offset(offSet).
		Limit(20).
		Find(&results).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("connection error. Retry later")
	}
	return &results, nil
}

func GetTodo(id, userID uint64) (*Todo, error) {
	todo := &Todo{}
	err := GetDB().Where("id = ? and user_id = ?", id, userID).First(todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("todo not found")
		}
		return nil, errors.New("connection error. Please retry")
	}
	return todo, nil
}

func (todo *Todo) UpdateTodo() error {
	//TODO: Get to-do by id and owner id
	//
	err := GetDB().Update(todo).Error
	if err != nil {
		return errors.New("update todo failed. Please retry")
	}
	return nil
}

func DeleteTodo(id uint64) error {
	err := GetDB().Delete(Todo{}, "id = ?", id).Error
	if err != nil {
		return errors.New("connection error. Please retry")
	}
	return nil
}
