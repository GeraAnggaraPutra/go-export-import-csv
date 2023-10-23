package crypt

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Generate bcrypt hash password.
func GenerateHashPassword(password string) (hashed string, err error) {
	crypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.Wrap(err, "error generate bcrypt hash password")
		return
	}

	hashed = string(crypt)

	return
}

// Compare bcrypt hashed password.
func CompareHashPassword(passwordInput, passwordDB string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordInput))
	return
}
