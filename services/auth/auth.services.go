package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"portfoyum-api/services/user"
	"portfoyum-api/utils"
	"portfoyum-api/utils/jwt"
	"portfoyum-api/utils/mail"
	_ "portfoyum-api/utils/mail"
	"portfoyum-api/utils/password"
)

func Signup(c *fiber.Ctx) error {
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

	data := new(TokenResponseDTO)
	//data.User = u.HttpFriendlyResponse()
	data.Token = jwt.Generate(t)

	return utils.Response(c, "User created", data)
}

func Login(c *fiber.Ctx) error {
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

	data := new(TokenResponseDTO)

	//data.User = u.HttpFriendlyResponse()
	//token := jwt.Generate(t)
	data.Token = jwt.Generate(t)

	//cookie := new(fiber.Cookie)
	//cookie.Name = "token"
	//cookie.Value = token
	//cookie.HTTPOnly = true
	//cookie.SameSite = "None"
	//cookie.Secure = false
	//cookie.Path = "/"
	//cookie.Expires = time.Now().AddDate(0, 1, 0)
	//
	//c.Cookie(cookie)

	return utils.Response(c, "Login successful", data)
}

func ForgotPassword(c *fiber.Ctx) error {
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

	return utils.Response(c, "Recovery email sent")
}

func verifyToken(token string) (*user.User, error) {
	payload, err := jwt.Verify(token)

	if err != nil {
		return nil, errors.New("token is not valid")
	}

	u := &user.User{}

	err = u.FindById(payload.ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return u, nil
}

func VerifyToken(c *fiber.Ctx) error {
	token := c.Params("token")

	if _, err := verifyToken(token); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return utils.Response(c, "Token is valid")
}

func ChangePassword(c *fiber.Ctx) error {
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

	return utils.Response(c, "Password changed")
}
