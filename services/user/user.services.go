package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"portfoyum/utils"
)

func getAuthorisedUser(c *fiber.Ctx) (*User, *fiber.Error) {
	id := utils.GetUserId(c)

	u := new(User)

	if id == nil {
		return u, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	err := u.FindById(*id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return u, nil
}

func UserMe(c *fiber.Ctx) error {
	u, err := getAuthorisedUser(c)
	if err != nil {
		return err
	}

	data := u.HttpFriendlyResponse()

	return c.JSON(utils.Response("Authorized user fetched", data))
}

func UserUpdate(c *fiber.Ctx) error {
	b := new(UserRequestDTO)

	u, err := getAuthorisedUser(c)
	if err != nil {
		return err
	}

	if err := utils.ParseBody(c, &b); err != nil {
		return err
	}

	if err := utils.Copy(u, b); err != nil {
		return err
	}

	if err := utils.Validate(u); err != nil {
		return  err
	}

	if err := u.Update(); err.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error.Error())
	}

	data := u.HttpFriendlyResponse()

	return c.JSON(utils.Response("User successfully updated", data))
}

func UserDelete(c *fiber.Ctx) error {
	u, err := getAuthorisedUser(c)
	if err != nil {
		return err
	}

	u.Active = false

	if err := u.Delete(); err.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error.Error())
	}

	return c.JSON(utils.Response("User successfully deleted"))
}
