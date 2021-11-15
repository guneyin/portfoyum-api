package admin

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"portfoyum/types"
	"portfoyum/utils"
	"portfoyum/utils/database"
	"portfoyum/utils/jwt"
	"portfoyum/utils/password"
)

func InitAdmin() {
	admin := new(Admin)
	var count int64

	database.DB.Model(admin).Count(&count)

	if count == 0 {
		pass := "Portfoyum++2021!"

		admin.Active = true
		admin.Email = "admin@portfoyum.com"
		admin.Password = password.Generate(pass)
		admin.Name = "portfoyum"
		admin.Surname = "Admin"

		if err := admin.Create(); err.Error != nil {
			color.Tag("warn").Printf("Error occured in InitAdmin\nError: %s", err.Error)
		}

		color.Tag("info").Printf("Admin user created: \nLogin: %s \nPass: %s", admin.Email, pass)
	}
}

func AdminLogin(c *fiber.Ctx) error {
	b := new(LoginRequestDTO)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	u := new(Admin)

	err := u.FindByEmail(b.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid email or password")
	}

	if err := password.Verify(u.Password, b.Password); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid email or password")
	}

	_ = utils.Copy(u, b)

	t := new(jwt.TokenPayload)
	t.ID = u.ID
	t.Email = u.Email
	t.Active = u.Active

	jwt := jwt.Generate(t)

	token := new(types.TToken)
	token.Token = jwt

	data := &LoginResponseDTO{
		Admin: u.HttpFriendlyResponse(),
		Token: token,
	}

	return c.JSON(utils.Response("Login successful", data))
}
