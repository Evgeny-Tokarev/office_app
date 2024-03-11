package token

import (
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/o1egl/paseto"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomString(10)
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)
	token, err := maker.CreateToken(username, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)

	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpirePasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomString(10), -time.Minute)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, util.ErrorTokenExpired.Error())
	require.Nil(t, payload)
}

func TestInvalidPasetoToken(t *testing.T) {
	payload, err := NewPayload(util.RandomString(10), time.Minute)
	require.NoError(t, err)
	key1 := util.RandomString(32)
	key2 := util.RandomString(32)

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(key1),
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	require.NoError(t, err)

	err = maker.paseto.Decrypt(token, []byte(key2), nil, nil)
	require.Error(t, err)
	require.EqualError(t, err, util.ErrInvalidTokenAuth.Error())
}
