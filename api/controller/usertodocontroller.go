package controller

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/model"
	"net/http"
	"strconv"
)

func GetUserTodo(c *gin.Context) {
	userID := c.Param("id")
	offset := c.Query("offset")
	uID, _ := strconv.ParseUint(userID, 10, 64)
	uOffSet, _ := strconv.ParseUint(offset, 10, 64)
	todos, err := model.GetUserTodos(uID, uOffSet)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Success: false,
		Error:   "",
		Data:    todos,
	})
}
