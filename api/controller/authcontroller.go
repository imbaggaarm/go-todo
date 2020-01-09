package controller

import (
	validator "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go-todo/api/model"
	"go-todo/api/util"
	"net/http"
)

func RegisterAccount(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err := user.Create()
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

func Login(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	user, err := model.Login(user.Email, user.Password)
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

//
//func Logout(c *gin.Context) {
//
//}

func ChangePassword(c *gin.Context) {

	email, ok := util.GetEmailFromContext(c)
	if !ok {
		c.JSON(http.StatusOK, model.UnauthorizedResponse())
		return
	}
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
	err = model.UpdatePassword(email, request.Password, request.NewPassword)
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

//
//func ForgotPassword(c *gin.Context) {
//
//}
//
//func ResetPassword(c *gin.Context) {
//
//}
