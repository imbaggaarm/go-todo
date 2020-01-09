package model

import (
	"errors"
	validator "github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"go-todo/api/auth"
	"go-todo/api/security"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	LastName    string    `json:"last_name"`
	FirstName   string    `json:"first_name"`
	Email       string    `json:"email" gorm:"unique;not null" valid:"email"`
	Password    string    `json:"password,omitempty" gorm:"not null" valid:"required,stringlength(6|256)"`
	Token       string    `json:"token" sql:"-"`
	NewPassword string    `json:"new_password,omitempty" sql:"-"`
}

func (user *User) Validate() error {

	if _, err := validator.ValidateStruct(user); err != nil {
		return err
	}
	temp := &User{}
	err := GetDB().Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("connection error")
	}
	if temp.Email != "" {
		return errors.New("email existed")
	}

	return nil
}

func (user *User) setToken() {
	user.Token = auth.CreateTokenString(user.ID, user.Email)
}

func (user *User) Create() error {
	if err := user.Validate(); err != nil {
		return err
	}

	hashedPassword, _ := security.Hash(user.Password)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return errors.New("connection error")
	}
	user.Password = "" // Remove password from response
	user.setToken()
	return nil
}

func verifyUser(email, password string) (*User, error) {
	user := &User{}
	err := GetDB().Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("account not found")
		}
		return nil, errors.New("connection error")
	}
	err = security.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.New("wrong password")
	}
	return user, nil
}

func Login(email, password string) (*User, error) {
	user, err := verifyUser(email, password)
	if err != nil {
		return nil, err
	}
	user.Password = "" // Remove password from response
	user.setToken()
	return user, nil
}

func UpdatePassword(email, password, newPassword string) error {
	user, err := verifyUser(email, password)
	if err != nil {
		return err
	}
	// Hash new password and update
	hashedPassword, _ := security.Hash(newPassword)
	GetDB().Model(user).Update("password", string(hashedPassword))

	// TODO: Handle logout of other devices
	return nil
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
func (user *User) Update() (*User, error) {
	currentUser, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	err = GetDB().Save(currentUser).Error
	if err != nil {
		return nil, errors.New("update todo failed. Please retry")
	}
	return currentUser, nil
}
