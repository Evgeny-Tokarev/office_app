package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"time"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < util.MinSecretKeySize {
		return nil, fmt.Errorf("invalid key size: should be at least %d characters", util.MinSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(id int64, role string, duration time.Duration) (string, error) {
	fmt.Println("creating jwt token")

	payload, err := NewPayload(id, role, duration)
	if err != nil {
		return "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(maker.secretKey))

}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(maker.secretKey), nil
	})
	if err != nil {
		var verr *jwt.ValidationError
		ok := errors.As(err, &verr)
		if ok && errors.Is(verr.Inner, util.ErrorTokenExpired) {
			return nil, util.ErrorTokenExpired
		}
		return nil, util.ErrorTokenInvalid
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, util.ErrorTokenInvalid
	}
	return payload, nil
}
