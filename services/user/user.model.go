package user

import (
	"github.com/imdario/mergo"
	"gorm.io/gorm"
	"log"
	"portfoyum-api/types"
	"portfoyum-api/utils"
	"portfoyum-api/utils/database"
)

type User struct {
	gorm.Model   `json:"-"`
	types.TEmail `gorm:"uniqueIndex"`
	types.TPassword
	types.TName
	types.TSurname
	types.TDateOfBirth
	types.TGender
	types.TPhone
}

func (u *User) Assign(v *UserRequestDTO) {
	if err := mergo.Merge(&u, &v); err != nil {
		log.Print(err.Error())
	}
}

func (u *User) HttpFriendlyResponse() *UserResponseDTO {
	r := new(UserResponseDTO)

	_ = utils.Copy(r, u)

	r.Password = ""

	return r
}

func (u *User) Create() *gorm.DB {
	return database.DB.Create(u)
}

func find(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&User{}).Take(dest, conds...)
}

func (u *User) Update() *gorm.DB {
	d := database.DB.Updates(u)

	if d.Error == nil {
		u.FindById(u.ID)
	}

	return d
}

func (u *User) Delete() *gorm.DB {
	return database.DB.Delete(u)
}

func (u *User) FindByEmail(email string) *gorm.DB {
	return find(u, "email = ?", email)
}

func (u *User) FindById(id uint) *gorm.DB {
	return find(u, "id = ?", id)
}
