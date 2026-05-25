package cryptoutil

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// ErrInvalidPassword is returned when the provided plain text password
// does not match the stored bcrypt hash during authentication.
var ErrInvalidPassword = errors.New("invalid password")

// PasswordHash converts a plain text password into a secure bcrypt hash.
// It automatically handles string-to-byte conversion and uses bcrypt.DefaultCost (10),
// which balances cryptographic security with server performance.
func PasswordHash(password string) (string, error) {
	// GenerateFromPassword hashes the password using the specified work factor (cost).
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	// Convert the byte slice back to a string for easier database storage.
	return string(hash), nil
}

// PasswordCheck compares a plain text password against an existing bcrypt hash.
// It returns nil if the password is correct, ErrInvalidPassword if it is incorrect,
// or any other internal error encountered during the comparison.
func PasswordCheck(password, hash string) error {
	// CompareHashAndPassword safely compares the hash and plaintext password to prevent timing attacks.
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	); err != nil {
		// Catch the specific bcrypt mismatch error and mask it with our custom error
		// to decouple the upper layers from the bcrypt package.
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrInvalidPassword
		}
		return err
	}

	return nil
}
