package mail

import (
	"errors"
	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"github.com/matcornic/hermes/v2"
	"net/mail"
	"portfoyum-api/config"
	"portfoyum-api/services/user"
)

type IMail interface {
	Email(user *user.User) hermes.Email
	Options() SendOptions
}

// SendOptions are options for sending an email
type SendOptions struct {
	To      string
	Subject string
}

type smtpAuthentication struct {
	Server         string
	Port           int
	SenderEmail    string
	SenderIdentity string
	SMTPUser       string
	SMTPPassword   string
}

func SendMail(e IMail, u *user.User) *fiber.Error {
	if config.Settings.Application.SendMails == false {
		return fiber.NewError(fiber.StatusInternalServerError, "Mail sender disabled by default")
	}

	h := hermes.Hermes{
		Product: hermes.Product{
			Name:        config.Settings.Application.Name,
			Link:        config.Settings.Application.Link,
			Logo:        "https://github.com/matcornic/hermes/blob/master/examples/gopher.png?raw=true",
			Copyright:   "portfoyum.com © 2020 - Tüm Hakları Saklıdır",
			TroubleText: "{ACTION} düğmesiyle ilgili sorun yaşıyorsanız, aşağıdaki URL'yi kopyalayıp web tarayıcınıza yapıştırın.",
		},
	}

	h.Theme = new(hermes.Flat)

	m := e.Email(u)
	m.Body.Name = u.Name + " " + u.Surname
	m.Body.Greeting = "Merhaba"
	m.Body.Signature = "Teşekkürler"

	o := e.Options()
	o.To = u.Email

	//u.Active = false

	htmlBytes, err := h.GenerateHTML(m)
	txtBytes, err := h.GeneratePlainText(m)

	err = send(o, htmlBytes, txtBytes)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

// send sends the email
func send(options SendOptions, htmlBody string, txtBody string) error {
	smtpConfig := smtpAuthentication{
		Server:         config.Settings.Mail.SmtpServer,
		Port:           config.Settings.Mail.SmtpPort,
		SenderEmail:    config.Settings.Mail.SenderEmail,
		SenderIdentity: config.Settings.Mail.SenderIdentity,
		SMTPPassword:   config.Settings.Mail.SmtpPassword,
		SMTPUser:       config.Settings.Mail.SmtpUserName,
	}

	if smtpConfig.Server == "" {
		return errors.New("SMTP server config is empty")
	}
	if smtpConfig.Port == 0 {
		return errors.New("SMTP port config is empty")
	}

	if smtpConfig.SMTPUser == "" {
		return errors.New("SMTP user is empty")
	}

	if smtpConfig.SenderIdentity == "" {
		return errors.New("SMTP sender identity is empty")
	}

	if smtpConfig.SenderEmail == "" {
		return errors.New("SMTP sender email is empty")
	}

	if options.To == "" {
		return errors.New("no receiver emails configured")
	}

	from := mail.Address{
		Name:    smtpConfig.SenderIdentity,
		Address: smtpConfig.SenderEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from.String())
	m.SetHeader("To", options.To)
	m.SetHeader("Subject", options.Subject)

	m.SetBody("text/plain", txtBody)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(smtpConfig.Server, smtpConfig.Port, smtpConfig.SMTPUser, smtpConfig.SMTPPassword)

	return d.DialAndSend(m)
}
