package hash

import "golang.org/x/crypto/bcrypt"

type PasswordHasher struct{}

func (p PasswordHasher) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
