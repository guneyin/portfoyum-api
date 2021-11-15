package auth

import (
	"portfoyum/services/user"
	"portfoyum/types"
)

type SignupRequestDTO struct {
	types.TEmail
	types.TPassword
	types.TName
	types.TSurname
}

type SignupResponseDTO struct {
	User *user.UserResponseDTO`json:"user"`
	Token *types.TToken
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
