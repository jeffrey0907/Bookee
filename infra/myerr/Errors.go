package myerr

import "errors"

var (
	ErrIllegalArgument = errors.New("ErrIllegalArgument")
	ErrNoUidINJWT      = errors.New("ErrNoUidInJWT")
)
