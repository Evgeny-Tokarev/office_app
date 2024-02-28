package util

import "errors"

const TimeLayout = "2006-01-02 15:04:05 -0700 MST"
const MinSecretKeySize = 32

var ErrorTokenExpired = errors.New("token has expired")
var ErrorTokenInvalid = errors.New("token is invalid")
var ErrInvalidTokenAuth = errors.New("invalid token authentication")
