package model

type UpdatePasswordRequest struct {
	Password    string `json:"password" valid:"required,stringlength(6|256)"`
	NewPassword string `json:"new_password" valid:"required,stringlength(6|256)"`
}
