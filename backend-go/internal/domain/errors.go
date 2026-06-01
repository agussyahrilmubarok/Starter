package domain

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrAlreadyExists  = errors.New("already exists")
	ErrInternalServer = errors.New("internal server error")

	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrTokenExpired = errors.New("token expired")
	ErrTokenInvalid = errors.New("token invalid")

	ErrInvalidInput    = errors.New("invalid input")
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrInvalidPassword = errors.New("password must be at least 8 characters")

	ErrEmailAlreadyUsed = errors.New("email already used")
	ErrWrongPassword    = errors.New("wrong password")
	ErrUserInactive     = errors.New("user is inactive")
)
