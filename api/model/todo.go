package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Todo struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	CompletedAt *time.Time `json:"completed_at" gorm:"index"`
	IsDone      bool       `json:"is_done" gorm:"default:0"`
	UserID      uint       `json:"user_id" gorm:"not null;index"`
	Note        string     `json:"note" gorm:"type:text"`
}

func (todo *Todo) CreateTodo() error {
	GetDB().Create(todo)
	if todo.ID <= 0 {
		return errors.New("create todo failed. Please retry")
	}
	return nil
}

func GetUserTodos(userID, offSet uint) (*[]Todo, error) {
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

func GetTodo(id, userID uint) (*Todo, error) {
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

func (todo *Todo) UpdateTodo() (*Todo, error) {

	currentTodo, err := GetTodo(todo.ID, todo.UserID)
	if err != nil {
		return nil, err
	}

	currentTodo.Title = todo.Title
	if todo.IsDone {
		current := time.Now()
		currentTodo.CompletedAt = &current
	}
	currentTodo.IsDone = todo.IsDone
	currentTodo.Note = todo.Note

	err = GetDB().Save(currentTodo).Error
	if err != nil {
		return nil, errors.New("update todo failed. Please retry")
	}
	return currentTodo, nil
}

func DeleteTodo(id uint, userID uint) error {
	err := GetDB().Delete(Todo{}, "id = ? and user_id = ?", id, userID).Error
	if err != nil {
		return errors.New("connection error. Please retry")
	}
	return nil
}
