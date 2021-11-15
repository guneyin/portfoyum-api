package mail

import (
	"github.com/matcornic/hermes/v2"
	"portfoyum-api/config"
	"portfoyum-api/services/user"
	"portfoyum-api/utils/jwt"
)

type Reset struct {}

func (r *Reset) Options() SendOptions {
	return SendOptions{
		Subject: "Şifrenizi sıfırlayın",
	}
}

func (r *Reset) Email(user *user.User) hermes.Email {
	t := new(jwt.TokenPayload)
	t.ID = user.ID
	t.Email = user.Email
	t.Active = user.Active

	token := jwt.Generate(t, "30m")

	return hermes.Email{
		Body: hermes.Body{
			Intros: []string{
				"Bu e-postayı, portfoyum hesabı için bir şifre sıfırlama talebi alındığı için aldınız.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Şifrenizi sıfırlamak için aşağıdaki düğmeyi tıklayın:",
					Button: hermes.Button{
						Color: "#DC4D2F",
						Text:  "Şifrenizi sıfırlayın",
						Link:  config.Settings.Application.Link + "/auth/password/forgot/" + token,
					},
				},
			},
			Outros: []string{
				"Parola sıfırlama talebinde bulunmadıysanız, başka bir işlem yapmanız gerekmez.",
			},
		},
	}
}