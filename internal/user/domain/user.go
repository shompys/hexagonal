package domain

import (
	"errors"

	"github.com/shompys/hexagonal/pkg/validation"
)

type User struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	UserName     string
	PasswordHash string
}

const (
	ID           = "ID"
	FirstName    = "FirstName"
	LastName     = "LastName"
	Email        = "Email"
	UserName     = "UserName"
	PasswordHash = "PasswordHash"
)

func NewInstance(id, firstName, lastName, email, userName, passwordHash string) (*User, error) {

	fields := map[string]string{
		ID:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		UserName:     userName,
		PasswordHash: passwordHash,
	}

	for fieldName, value := range fields {
		if err := validation.ValidateStringNotEmpty(fieldName, value); err != nil {
			return nil, errors.New(err.Error())
		}
	}

	if ok := validation.IsEmailValid(email); !ok {
		return nil, errors.New("invalid email format")
	}

	return &User{
		ID:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		UserName:     userName,
		PasswordHash: passwordHash,
	}, nil
}
