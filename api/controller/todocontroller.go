package controller

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/model"
	u "go-todo/api/util"
	"net/http"
	"strconv"
)

func CreateTodo(c *gin.Context) {
	todo := &model.Todo{}
	if err := c.BindJSON(todo); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	userID, ok := u.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusOK, model.UnauthorizedResponse())
		return
	}
	todo.UserID = userID
	err := todo.CreateTodo()
	if err != nil {
		c.JSON(http.StatusOK, model.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse(todo))
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	uID, _ := strconv.ParseUint(id, 10, 64)

	userID, ok := u.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusOK, model.UnauthorizedResponse())
		return
	}
	todo, err := model.GetTodo(uint(uID), userID)
	if err != nil {
		c.JSON(http.StatusOK, model.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse(todo))
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

	userID, ok := u.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusOK, model.UnauthorizedResponse())
		return
	}
	todo.UserID = userID
	updatedTodo, err := todo.UpdateTodo()
	if err != nil {
		c.JSON(http.StatusOK, model.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse(updatedTodo))
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	uID, _ := strconv.ParseUint(id, 10, 10)

	userID, ok := u.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusOK, model.UnauthorizedResponse())
		return
	}
	err := model.DeleteTodo(uint(uID), userID)
	if err != nil {
		c.JSON(http.StatusOK, model.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}
