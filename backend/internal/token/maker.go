package token

import "time"

var makerMap = map[string]func(secret string) (Maker, error){
	"jwt": func(secret string) (Maker, error) {
		return NewJWTMaker(secret)
	},
	"paseto": func(secret string) (Maker, error) {
		return NewPasetoMaker(secret)
	},
}

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

func NewMaker(mt string, secret string) (Maker, error) {
	return makerMap[mt](secret)
}
