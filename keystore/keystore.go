package keystore

import "errors"

var (
	ErrNotSetPassword = errors.New("Password can not be set")
	ErrPassordLenOver = errors.New("Password len 32 is over")
)
