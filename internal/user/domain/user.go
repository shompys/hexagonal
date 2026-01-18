package domain

import (
	"errors"
	"time"

	"github.com/shompys/hexagonal/pkg/validation"
)

type User struct {
	id           UserIDVO
	firstName    string
	lastName     string
	email        string
	userName     string
	passwordHash UserPasswordVO //debe pasarse sin puntero porque esto asegura que es valido por ende no pueden poner nil
	createdAt    time.Time
	updatedAt    time.Time
}

const (
	ID           = "ID"
	FirstName    = "FirstName"
	LastName     = "LastName"
	Email        = "Email"
	UserName     = "UserName"
	PasswordHash = "PasswordHash"
)

func NewUser(firstName, lastName, email, userName string, passwordHash UserPasswordVO) (*User, error) {

	fields := map[string]string{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UserName:  userName,
	}

	for fieldName, value := range fields {
		if err := validation.ValidateStringNotEmpty(fieldName, value); err != nil {
			return nil, err
		}
	}

	if ok := validation.IsEmailValid(email); !ok {
		return nil, errors.New("invalid email format")
	}

	now := time.Now()

	return &User{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		userName:     userName,
		passwordHash: passwordHash,
		createdAt:    now,
		updatedAt:    now,
	}, nil
}

func RestoreUser(id UserIDVO, firstName, lastName, email, userName string, passwordHash UserPasswordVO, createdAt, updatedAt time.Time) *User {
	return &User{
		id:           id,
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		userName:     userName,
		passwordHash: passwordHash,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
}

func (u *User) ID() string           { return u.id.Value() }
func (u *User) FirstName() string    { return u.firstName }
func (u *User) LastName() string     { return u.lastName }
func (u *User) Email() string        { return u.email }
func (u *User) UserName() string     { return u.userName }
func (u *User) PasswordHash() string { return u.passwordHash.Value() }
func (u *User) CreatedAt() time.Time { return u.createdAt }
func (u *User) UpdatedAt() time.Time { return u.updatedAt }

func (u *User) SetID(id UserIDVO) {
	u.id = id
}

func (u *User) SetFirstName(firstName string) error {
	if err := validation.ValidateStringNotEmpty(FirstName, firstName); err != nil {
		return err
	}
	u.firstName = firstName
	u.updatedAt = time.Now()
	return nil
}
func (u *User) SetLastName(lastName string) error {
	if err := validation.ValidateStringNotEmpty(LastName, lastName); err != nil {
		return err
	}
	u.lastName = lastName
	u.updatedAt = time.Now()
	return nil
}
func (u *User) SetEmail(email string) error {
	if err := validation.ValidateStringNotEmpty(Email, email); err != nil {
		return err
	}
	if ok := validation.IsEmailValid(email); !ok {
		return errors.New("invalid email format")
	}
	u.email = email
	u.updatedAt = time.Now()
	return nil
}
func (u *User) SetUserName(userName string) error {
	if err := validation.ValidateStringNotEmpty(UserName, userName); err != nil {
		return err
	}
	u.userName = userName
	u.updatedAt = time.Now()
	return nil
}

func (u *User) SetPasswordHash(passwordHash UserPasswordVO) {
	u.passwordHash = passwordHash
	u.updatedAt = time.Now()
}
