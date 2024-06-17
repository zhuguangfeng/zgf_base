package errs

import (
	"errors"
)

var (
	ErrDuplicatePhone  = errors.New("手机号码冲突")
	ErrUserNotFound    = errors.New("record not found")
	ErrInvalidPassword = errors.New("密码错误")
)
