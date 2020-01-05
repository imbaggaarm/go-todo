package model

type UpdatePasswordRequest struct {
	Email       string `json:"email" valid:"required,email"`
	Password    string `json:"password" valid:"required,stringlength(6|256)"`
	NewPassword string `json:"new_password" valid:"required,stringlength(6|256)"`
}


