package user

import "portfoyum-api/types"

type UserRequestDTO struct {
	types.TEmail
	types.TPassword
	types.TName
	types.TSurname
	types.TDateOfBirth
	types.TGender
	types.TPhone
}

type UserResponseDTO struct {
	types.TUID
	UserRequestDTO
}
