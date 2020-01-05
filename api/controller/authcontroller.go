package controller

import (
	validator "github.com/asaskevich/govalidator"
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
	request := &model.UpdatePasswordRequest{}
	if err := c.BindJSON(request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	_, err := validator.ValidateStruct(request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}
	// Update password
	resp := model.UpdatePassword(request.Email, request.Password, request.NewPassword)
	c.JSON(http.StatusOK, resp)
}

func ForgotPassword(c *gin.Context) {

}

func ResetPassword(c *gin.Context) {

}
