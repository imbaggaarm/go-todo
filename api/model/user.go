package model

import (
	"errors"
	validator "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"go-todo/api/auth"
	"go-todo/api/security"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	Email       string `json:"email" gorm:"unique;not null" valid:"email"`
	Username    string `json:"username" gorm:"unique;not null"`
	Password    string `json:"password,omitempty" gorm:"not null"`
	Token       string `sql:"-"`
	NewPassword string `json:"new_password,omitempty" sql:"-"`
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

func (user *User) setToken() {
	user.Token = auth.CreateTokenString(user.ID)
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

func verifyUser(email, password string) (*User, *Response) {
	user := &User{}
	err := GetDB().Where("email = ? or username = ?", email, email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &Response{
				Success: false,
				Error:   "Account not found",
				Data:    nil,
			}
		}
		return nil, &Response{
			Success: false,
			Error:   "Connection error. Please retry",
			Data:    nil,
		}
	}
	err = security.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, &Response{
			Success: false,
			Error:   "Wrong password",
			Data:    nil,
		}
	}
	return user, nil
}

func Login(email, password string) Response {
	user, resp := verifyUser(email, password)
	if resp != nil {
		return *resp
	}
	user.Password = "" // Remove password from response
	user.setToken()
	return Response{
		Success: true,
		Error:   "",
		Data:    user,
	}
}

func UpdatePassword(email, password, newPassword string) Response {
	user, resp := verifyUser(email, password)
	if resp != nil {
		return *resp
	}
	// Hash new password and update
	hashedPassword, _ := security.Hash(newPassword)
	GetDB().Model(user).Update("password", string(hashedPassword))

	// TODO: Handle logout of other devices
	return Response{
		Success: true,
		Error:   "",
		Data:    nil,
	}
}

func GetUser(userID uint) (*User, error) {
	user := &User{}
	err := GetDB().First(user, userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("account not found")
		}
		return nil, errors.New("connection error. Please retry")
	}
	return user, nil
}