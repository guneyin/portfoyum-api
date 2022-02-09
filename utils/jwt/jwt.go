package jwt

import (
	"errors"
	"fmt"
	"portfoyum-api/config"
	"portfoyum-api/types"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenPayload struct {
	ID uint `json:"id"`
	types.TEmail
}

type Claim struct {
	TokenPayload
	jwt.StandardClaims
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

	claim := &Claim{
		TokenPayload: *payload,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(v).Unix(),
			Id:        string(payload.ID),
			IssuedAt:  time.Now().Unix(),
			Issuer:    payload.Email,
			NotBefore: 0,
			Subject:   "",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"exp":  time.Now().Add(v).Unix(),
	//	"user": payload,
	//})

	token, err := t.SignedString([]byte(config.Settings.Jwt.TokenKey))

	if err != nil {
		panic(err)
	}

	return token
}

func parse(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Settings.Jwt.TokenKey), nil
	})

	//return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
	//	// Don't forget to validate the alg is what you expect:
	//	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
	//	}
	//
	//	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	//	return []byte(config.Settings.Jwt.TokenKey), nil
	//})
}

// Verify verifies the jwt token against the secret
func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)

	if err != nil {
		return nil, err
	}

	//claims, ok := parsed.Claims.(jwt.MapClaims)
	claims, ok := parsed.Claims.(*Claim)

	if !ok {
		return nil, err
	}

	// Getting ID, it's an interface{} so I need to cast it to uint
	//id, ok := claims["id"].(float64)
	//email, ok := claims["email"].(string)
	//uid, err := strconv.ParseUint(id, 10, 32)
	//uid, err := strconv.Atoi(id)
	//uid, err := uuid.FromString(id)

	if !ok {
		return nil, errors.New("something went wrong")
	}

	t := new(TokenPayload)
	t.ID = claims.ID
	t.Email = claims.Email

	return t, nil
}
