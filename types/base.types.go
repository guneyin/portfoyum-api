package types

import (
	"github.com/gofrs/uuid"
	"time"
)

type TUID struct {
	UID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

type TActive struct {
	Active bool
}

type TEmail struct {
	Email string `gorm:"uniqueIndex;not null;size:50" json:"email" validate:"required,email"`
}

type TPassword struct {
	Password string `gorm:"not null;size:150" json:"password,omitempty" validate:"required,password"`
}

type TName struct {
	Name string `gorm:"unique;not null;size:50" json:"name" validate:"required,min=3"`
}

type TSurname struct {
	Surname string `gorm:"not null;size:50" json:"surname" validate:"required,min=3"`
}

type TToken struct {
	Token string `json:"token"`
}

type TDateOfBirth struct {
	DateOfBirth time.Time `json:"date_of_birth" validate:"omitempty"`
}

type TGender struct {
	Gender string `gorm:"size:1" json:"gender" validate:"omitempty,oneof=m f"`
}

type TPhone struct {
	Phone string `gorm:"size:20" json:"phone" validate:"omitempty"`
}
