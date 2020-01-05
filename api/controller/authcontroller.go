package controller

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/model"
	"net/http"
)

func RegisterAccount(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	resp := user.Create()
	c.JSON(http.StatusOK, resp)
}

func Login(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	resp := model.Login(user.Email, user.Password)
	c.JSON(http.StatusOK, resp)
}

func Logout(c *gin.Context) {

}

func ChangePassword(c *gin.Context) {

}

func ForgotPassword(c *gin.Context) {

}

func ResetPassword(c *gin.Context) {

}
