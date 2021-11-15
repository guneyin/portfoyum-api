package auth

import (
	"errors"
	"portfoyum/services/user"
	"portfoyum/types"
	"portfoyum/utils"
	"portfoyum/utils/jwt"
	"portfoyum/utils/mail"
	_ "portfoyum/utils/mail"
	"portfoyum/utils/password"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthSignup(c *fiber.Ctx) error {
	b := new(SignupRequestDTO)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	u := new(user.User)

	err := u.FindByEmail(b.Email).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusBadRequest, "Email already exists")
	}

	_ = utils.Copy(u, b)

	u.Password = password.Generate(b.Password)

	if err := u.Create(); err.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error.Error())
	}

	t := new(jwt.TokenPayload)
	t.ID = u.ID
	t.Email = u.Email
	t.Active = u.Active

	a := new(types.TToken)
	a.Token = jwt.Generate(t)

	data := &SignupResponseDTO{
		User: u.HttpFriendlyResponse(),
		Token: a,
	}

	return c.JSON(utils.Response("User created", data))
}

func AuthLogin(c *fiber.Ctx) error {
	b := new(LoginRequestDTO)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	u := new(user.User)

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

	data := &SignupResponseDTO{
		User: u.HttpFriendlyResponse(),
		Token: token,
	}

	return c.JSON(utils.Response("Login successful", data))
}

func AuthForgotPassword(c *fiber.Ctx) error {
	b := new(PasswordForgotRequestDTO)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Message)
	}

	u := user.User{}

	err := u.FindByEmail(b.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusBadRequest, "User not found")
	}

	if err := mail.SendMail(new(mail.Reset), &u); err != nil {
		return err
	}

	return c.JSON(utils.Response("Recovery email sent"))
}

func verifyToken(token string) (*user.User, error) {
	payload, err := jwt.Verify(token)

	if err != nil {
		return nil, errors.New("Token is not valid")
	}

	u := &user.User{}

	err = u.FindById(payload.ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("User not found")
	}

	return u, nil
}

func AuthVerifyToken(c *fiber.Ctx) error {
	token := c.Params("token")

	if _, err := verifyToken(token); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(utils.Response("Token is valid"))
}

func AuthChangePassword(c *fiber.Ctx) error {
	b := new(PasswordChangeRequestDTO)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	u, err := verifyToken(b.Token)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if b.Password1 != b.Password2 {
		return fiber.NewError(fiber.StatusBadRequest, "Passwords not identical")
	}

	u.Password = password.Generate(b.Password1)

	if err := u.Update(); err.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error.Error())
	}

	return c.JSON(utils.Response("Password changed"))
}