package authentication

import (
	"sthl/constants"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", constants.ErrBadRequest
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), constants.UserPwHashCost)
	return string(bytes), err
}

func CompareHashPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
