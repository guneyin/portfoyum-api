package admin

import (
	"portfoyum/types"
)

type AdminRequestDTO struct {
	types.TEmail
	types.TPassword
	types.TName
	types.TSurname
}

type AdminResponseDTO struct {
	types.TID
	AdminRequestDTO
}

type LoginRequestDTO struct {
	types.TEmail
	types.TPassword
}

type LoginResponseDTO struct {
	Admin *AdminResponseDTO`json:"user"`
	Token *types.TToken
}