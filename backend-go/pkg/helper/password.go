package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PasswordEncrypt(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[ERROR] failed to encrypt password: %v\n", err)
	}
	return string(hashed)
}

func PasswordCompare(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}
