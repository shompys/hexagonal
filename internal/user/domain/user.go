package domain

import (
	"errors"

	"github.com/shompys/hexagonal/pkg/validation"
)

type User struct {
	id           string
	firstName    string
	lastName     string
	email        string
	userName     string
	passwordHash UserPasswordVO //debe pasarse sin puntero porque esto asegura que es valido por ende no pueden poner nil
}

const (
	ID           = "ID"
	FirstName    = "FirstName"
	LastName     = "LastName"
	Email        = "Email"
	UserName     = "UserName"
	PasswordHash = "PasswordHash"
)

func NewInstance(id, firstName, lastName, email, userName string, passwordHash UserPasswordVO) (*User, error) {

	fields := map[string]string{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UserName:  userName,
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
		id,
		firstName,
		lastName,
		email,
		userName,
		passwordHash,
	}, nil
}

func (u *User) ID() string           { return u.id }
func (u *User) FirstName() string    { return u.firstName }
func (u *User) LastName() string     { return u.lastName }
func (u *User) Email() string        { return u.email }
func (u *User) UserName() string     { return u.userName }
func (u *User) PasswordHash() string { return u.passwordHash.Value() }

func (u *User) SetFirstName(firstName string) error {
	if err := validation.ValidateStringNotEmpty(FirstName, firstName); err != nil {
		return errors.New(err.Error())
	}
	u.firstName = firstName
	return nil
}
func (u *User) SetLastName(lastName string) error {
	if err := validation.ValidateStringNotEmpty(LastName, lastName); err != nil {
		return errors.New(err.Error())
	}
	u.lastName = lastName
	return nil
}
func (u *User) SetEmail(email string) error {
	if err := validation.ValidateStringNotEmpty(Email, email); err != nil {
		return errors.New(err.Error())
	}
	if ok := validation.IsEmailValid(email); !ok {
		return errors.New("invalid email format")
	}
	u.email = email
	return nil
}
func (u *User) SetUserName(userName string) error {
	if err := validation.ValidateStringNotEmpty(UserName, userName); err != nil {
		return errors.New(err.Error())
	}
	u.userName = userName
	return nil
}

func (u *User) SetPasswordHash(passwordHash UserPasswordVO) {
	u.passwordHash = passwordHash
}
