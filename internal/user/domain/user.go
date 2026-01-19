package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/shompys/hexagonal/pkg/validation"
)

type UserStatus string

const (
	StatusActive   UserStatus = "active"
	StatusInactive UserStatus = "inactive"
)

type StatusChanges struct {
	Status    UserStatus
	ChangedAt time.Time
}

type User struct {
	id            UserIDVO
	firstName     string
	lastName      string
	email         string
	userName      string
	passwordHash  UserPasswordVO //debe pasarse sin puntero porque esto asegura que es valido por ende no pueden poner nil
	createdAt     time.Time
	updatedAt     time.Time
	status        UserStatus
	statusHistory []StatusChanges
}

const (
	ID           = "ID"
	FirstName    = "FirstName"
	LastName     = "LastName"
	Email        = "Email"
	UserName     = "UserName"
	PasswordHash = "PasswordHash"
	Status       = "status"
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

	initialStatus := StatusChanges{
		Status:    StatusActive,
		ChangedAt: now,
	}

	return &User{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		userName:      userName,
		passwordHash:  passwordHash,
		createdAt:     now,
		updatedAt:     now,
		status:        StatusActive,
		statusHistory: []StatusChanges{initialStatus},
	}, nil
}

func RestoreUser(
	id UserIDVO,
	firstName,
	lastName,
	email,
	userName string,
	passwordHash UserPasswordVO,
	createdAt, updatedAt time.Time,
	status UserStatus,
	statusHistory []StatusChanges,
) *User {
	return &User{
		id:            id,
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		userName:      userName,
		passwordHash:  passwordHash,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		status:        status,
		statusHistory: statusHistory,
	}
}

func (u *User) ID() string                     { return u.id.Value() }
func (u *User) FirstName() string              { return u.firstName }
func (u *User) LastName() string               { return u.lastName }
func (u *User) Email() string                  { return u.email }
func (u *User) UserName() string               { return u.userName }
func (u *User) PasswordHash() string           { return u.passwordHash.Value() }
func (u *User) CreatedAt() time.Time           { return u.createdAt }
func (u *User) UpdatedAt() time.Time           { return u.updatedAt }
func (u *User) Status() UserStatus             { return u.status }
func (u *User) StatusHistory() []StatusChanges { return u.statusHistory }

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

func (u *User) setStatus(status UserStatus) error {

	if u.status == status {
		return fmt.Errorf("user status is already [%s]", u.status)
	}

	u.status = status

	now := time.Now()

	statusInsert := StatusChanges{
		Status:    status,
		ChangedAt: now,
	}

	u.statusHistory = append(u.statusHistory, statusInsert)
	return nil
}

func (u *User) Deactivate() error {
	return u.setStatus(StatusInactive)
}
func (u *User) Activate() error {
	return u.setStatus(StatusActive)
}
