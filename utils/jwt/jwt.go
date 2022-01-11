package jwt

import (
	"errors"
	"fmt"
	"portfoyum-api/config"
	"portfoyum-api/types"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type TokenPayload struct {
	types.TUID
	types.TEmail
	types.TActive
}

func Generate(payload *TokenPayload, duration ...string) string {
	var d string

	if len(duration) > 0 {
		d = duration[0]
	} else {
		d = config.Settings.Jwt.TokenExp
	}
	v, err := time.ParseDuration(d)

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(v).Unix(),
		"ID":     payload.UID,
		"email":  payload.Email,
		"active": payload.Active,
	})

	token, err := t.SignedString([]byte(config.Settings.Jwt.TokenKey))

	if err != nil {
		panic(err)
	}

	return token
}

func parse(token string) (*jwt.Token, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Settings.Jwt.TokenKey), nil
	})
}

// Verify verifies the jwt token against the secret
func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)

	if err != nil {
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	// Getting ID, it's an interface{} so I need to cast it to uint
	id, ok := claims["ID"].(string)
	uid, err := uuid.FromString(id)

	if (!ok) && (err != nil) {
		return nil, errors.New("something went wrong")
	}

	t := new(TokenPayload)
	t.UID = uid

	return t, nil
}
