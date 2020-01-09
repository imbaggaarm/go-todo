package controller

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/model"
	u "go-todo/api/util"
	"net/http"
	"strconv"
)

func GetUserTodo(c *gin.Context) {
	userID := c.Param("id")
	offset := c.Query("offset")
	uID, _ := strconv.ParseUint(userID, 10, 64)
	uOffSet, _ := strconv.ParseUint(offset, 10, 64)

	//TODO: Get user id from token to authenticate user whether user can get todo or not
	tkUserID, ok := u.GetUserIDFromContext(c)
	if !ok || tkUserID != uint(uID) {
		c.JSON(http.StatusOK, model.UnauthorizedResponse())
		return
	}

	todos, err := model.GetUserTodos(uint(uID), uint(uOffSet))
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
		Data:    todos,
	})
}
