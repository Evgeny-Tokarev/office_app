package util

import (
	"github.com/caarlos0/env/v6"
	"github.com/evgeny-tokarev/office_app/backend/internal/config"
	"time"
)

import (
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		return "", err
	}

	secret := cfg.JwtSecret
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
