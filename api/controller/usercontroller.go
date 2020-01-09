package controller

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/model"
	u "go-todo/api/util"
	"net/http"
	"strconv"
)

func GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	uID, _ := strconv.ParseUint(id, 10, 64)
	//TODO: Verify user_id from token
	userID, ok := u.GetUserIDFromContext(c)
	if !ok || userID != uint(uID) {
		c.JSON(http.StatusOK, model.UnauthorizedResponse())
		return
	}

	user, err := model.GetUser(uint(uID))
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Error:   "",
		Data:    user,
	})
}

func UpdateUserInfo(c *gin.Context) {
	id := c.Param("id")
	uID, _ := strconv.ParseUint(id, 10, 64)
	//TODO: Verify user_id from token
	userID, ok := u.GetUserIDFromContext(c)
	if !ok || userID != uint(uID) {
		c.JSON(http.StatusOK, model.UnauthorizedResponse())
		return
	}
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	user.ID = userID
	updatedUser, err := user.Update()
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}
	updatedUser.Password = ""
	c.JSON(http.StatusOK, model.Response{
		Success: true,
		Error:   "",
		Data:    updatedUser,
	})
}
