package domain

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserEmailUsed       = errors.New("user email is already used")
	ErrUserPasswordInvalid = errors.New("user password does not match")
)
