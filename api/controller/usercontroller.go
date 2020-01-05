package controller

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/model"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	uID, _ := strconv.ParseUint(id, 10, 64)
	user, err := model.GetUser(uint(uID))
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
		Data:    user,
	})
}

func UpdateUser(c *gin.Context) {

}
