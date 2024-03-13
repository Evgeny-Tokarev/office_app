package util

import "errors"

const TimeLayout = "2006-01-02 15:04:05 -0700 MST"
const MinSecretKeySize = 32

var ReqOwners = map[string]uint8{
	"user":      1,
	"moderator": 2,
	"admin":     3,
}

var ErrorTokenExpired = errors.New("token has expired")
var ErrorTokenInvalid = errors.New("token is invalid")
var ErrInvalidTokenAuth = errors.New("invalid token authentication")
var AuthorizationTypeBearer = "Bearer"
var AuthorizationType = "Authorization"
