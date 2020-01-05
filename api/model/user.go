package model

import (
	validator "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"go-todo/api/auth"
	"go-todo/api/security"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email" gorm:"unique;not null" valid:"email"`
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password,omitempty" gorm:"not null"`
	Token     string `sql:"-"`
}

func (user *User) Validate() (Response, bool) {

	if _, err := validator.ValidateStruct(user); err != nil {
		return Response{
			Success: false,
			Error:   "Invalid email address",
			Data:    nil,
		}, false
	}

	if len(user.Password) < 6 {
		return Response{
			Success: false,
			Error:   "Short password",
			Data:    nil,
		}, false
	}

	temp := &User{}

	err := GetDB().Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Response{
			Success: false,
			Error:   "Connection error. Please retry.",
			Data:    nil,
		}, false
	}
	if temp.Email != "" {
		return Response{
			Success: false,
			Error:   "Email existed",
			Data:    nil,
		}, false
	}

	return Response{
		Success: true,
		Error:   "",
		Data:    nil,
	}, true
}

func (user *User) Create() Response {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := security.Hash(user.Password)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return Response{
			Success: false,
			Error:   "Failed to create account, connection error.",
			Data:    nil,
		}
	}
	user.Password = "" // Remove password from response
	user.setToken()
	return Response{
		Success: true,
		Error:   "",
		Data:    user,
	}
}

func Login(email, password string) Response {
	user := &User{}
	err := GetDB().Where("email = ? or username = ?", email, email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Response{
				Success: false,
				Error:   "Account not found",
				Data:    nil,
			}
		}
		return Response{
			Success: false,
			Error:   "Connection error. Please retry",
			Data:    nil,
		}
	}
	err = security.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return Response{
			Success: false,
			Error:   "Wrong password",
			Data:    nil,
		}
	}
	user.Password = "" // Remove password from response
	user.setToken()
	return Response{
		Success: true,
		Error:   "",
		Data:    user,
	}
}

func (user *User) setToken() {
	user.Token = auth.CreateTokenString(user.ID)
}
