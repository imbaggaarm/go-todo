package controller

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/model"
	"net/http"
	"strconv"
)

func CreateTodo(c *gin.Context) {
	todo := &model.Todo{}
	if err := c.BindJSON(todo); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err := todo.CreateTodo()
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Error:   "",
		Data:    todo,
	})
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	uID, _ := strconv.ParseUint(id, 10, 64)

	// TODO: Get user id from token
	todo, err := model.GetTodo(uID, 0)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Error:   "",
		Data:    todo,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	uID, _ := strconv.ParseUint(id, 10, 64)
	todo := &model.Todo{}
	if err := c.BindJSON(todo); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	todo.ID = uint(uID)
	// TODO: Get user id from token
	err := todo.UpdateTodo()
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Error:   "",
		Data:    todo,
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	uID, _ := strconv.ParseUint(id, 10, 10)
	err := model.DeleteTodo(uID)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Error:   "",
		Data:    nil,
	})
}
