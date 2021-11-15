package admin

import (
	"github.com/gofrs/uuid"
	"github.com/imdario/mergo"
	"gorm.io/gorm"
	"log"
	"portfoyum/types"
	"portfoyum/utils"
	"portfoyum/utils/database"
)

type Admin struct {
	types.TBaseModel `json:"-"`
	types.TEmail
	types.TPassword
	types.TName
	types.TSurname
}

func (a *Admin) Assign(v *Admin) {
	if err := mergo.Merge(&a, &v); err != nil {
		log.Print(err.Error())
	}
}

func (a *Admin) HttpFriendlyResponse() *AdminResponseDTO {
	r := new(AdminResponseDTO)

	_ = utils.Copy(r, a)

	r.Password = ""

	return r
}

func (a *Admin) Create() *gorm.DB {
	return database.DB.Create(a)
}

func find(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&Admin{}).Take(dest, conds...)
}

func (a *Admin) Update() *gorm.DB {
	d := database.DB.Updates(a)

	if d.Error == nil {
		a.FindById(a.ID)
	}

	return d
}

func (a *Admin) Delete() *gorm.DB {
	return database.DB.Delete(a)
}

func (a *Admin) FindByEmail(email string) *gorm.DB {
	return find(a, "email = ?", email)
}

func (a *Admin) FindById(id uuid.UUID) *gorm.DB {
	return find(a, "id = ?", id)
}

