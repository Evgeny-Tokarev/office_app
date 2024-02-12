package util

import (
	"errors"
	"net/http/httptest"
)

type CustomResponseRecorder struct {
	*httptest.ResponseRecorder
	encodingError error
	useCustom     bool
}

func NewCustomResponseRecorder() *CustomResponseRecorder {
	return &CustomResponseRecorder{
		ResponseRecorder: httptest.NewRecorder(),
		encodingError:    errors.New("encoding failure"),
		useCustom:        true,
	}
}

func (cr *CustomResponseRecorder) Write(b []byte) (int, error) {
	if cr.useCustom {
		cr.useCustom = false
		return 0, cr.encodingError
	}
	return cr.ResponseRecorder.Write(b)
}
