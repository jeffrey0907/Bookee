package myerr

import "errors"

var (
    ErrIllegalArgument = errors.New("IllegalArgument")
    ErrNoUidINJWT      = errors.New("NoUidInJWT")
    ErrNotExist        = errors.New("NotExtsts")
)
