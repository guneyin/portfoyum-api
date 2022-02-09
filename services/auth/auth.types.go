package auth

import (
	"portfoyum-api/types"
)

type SignupRequestDTO struct {
	types.TEmail
	types.TPassword
	types.TName
	types.TSurname
}

type TokenResponseDTO struct {
	//User *user.UserResponseDTO `json:"user"`
	types.TToken
}

type LoginRequestDTO struct {
	types.TEmail
	types.TPassword
}

type PasswordForgotRequestDTO struct {
	types.TEmail
}

type PasswordChangeRequestDTO struct {
	types.TToken
	Password1 string `json:"password1,omitempty" validate:"required,password"`
	Password2 string `json:"password2,omitempty" validate:"required,password"`
}
