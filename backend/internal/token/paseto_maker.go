package token

import (
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) < chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: should be at least %d characters", chacha20poly1305.KeySize)
	}
	if len(symmetricKey) > chacha20poly1305.KeySize {
		symmetricKey = symmetricKey[:32]
	}
	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey)}, nil
}

func (maker *PasetoMaker) CreateToken(id int64, role string, duration time.Duration) (string, error) {
	payload, err := NewPayload(id, role, duration)
	if err != nil {
		return "", nil
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)

}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, util.ErrorTokenInvalid
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
