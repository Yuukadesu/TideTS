package authority

import (
	commons "github.com/hanami/tidets/commons/errors"

	"golang.org/x/crypto/bcrypt"
)

const maxPasswordBytes = 72

func hashPassword(password string) (string, error) {
	if password == "" {
		return "", commons.ErrAuthPasswordRequired
	}
	if len(password) > maxPasswordBytes {
		return "", commons.ErrAuthPasswordTooLong
	}
	sum, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", commons.Wrap("auth", commons.CodeInternal, "hash password", err)
	}
	return string(sum), nil
}

func passwordMatches(hash, password string) bool {
	if hash == "" || password == "" {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
