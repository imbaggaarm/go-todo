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
		c.JSON(http.StatusOK, util.ErrorResponse(err))
		return
	}
	user.Password = "" // Remove password from response
	c.JSON(http.StatusOK, util.SuccessResponse(user))
}

func Login(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	user, err := model.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusOK, util.ErrorResponse(err))
		return
	}
	user.Password = "" // Remove password from response
	c.JSON(http.StatusOK, util.SuccessResponse(user))
}

//
//func Logout(c *gin.Context) {
//
//}

func ChangePassword(c *gin.Context) {
	email, ok := util.GetEmailFromContext(c)
	if !ok {
		c.JSON(http.StatusOK, util.UnauthorizedResponse())
		return
	}
	request := &util.UpdatePasswordRequest{}
	if err := c.BindJSON(request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	_, err := validator.ValidateStruct(request)
	if err != nil {
		c.JSON(http.StatusOK, util.ErrorResponse(err))
		return
	}
	// Update password
	err = model.UpdatePassword(email, request.Password, request.NewPassword)
	if err != nil {
		c.JSON(http.StatusOK, util.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, util.SuccessResponse(nil))
}

//
//func ForgotPassword(c *gin.Context) {
//
//}
//
//func ResetPassword(c *gin.Context) {
//
//}
