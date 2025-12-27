package domain

import (
	"errors"

	"github.com/shompys/hexagonal/pkg/validation"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
}

type UserPasswordVO struct {
	value string // El hash resultante
}

// NewUserPassword es el constructor del Value Object.
// Fíjate que obligamos a pasar un PasswordHasher (el puerto).
func NewUserPassword(rawPassword string, hasher PasswordHasher) (UserPasswordVO, error) {

	if err := validation.ValidateStringNotEmpty(PasswordHash, rawPassword); err != nil {
		return UserPasswordVO{}, errors.New(err.Error())
	}
	// 1. REGLA DE NEGOCIO: Validamos la contraseña ANTES de hashearla
	if len(rawPassword) < 8 {
		return UserPasswordVO{}, errors.New("la contraseña debe tener al menos 8 caracteres")
	}

	// 2. USO DEL PUERTO: Hasheamos usando la interfaz
	hashed, err := hasher.HashPassword(rawPassword)
	if err != nil {
		return UserPasswordVO{}, err
	}

	return UserPasswordVO{value: hashed}, nil
}

// RestoreUserPassword permite crear el VO a partir de un hash existente (ej. desde la DB o para mocks)
func RestoreUserPassword(hash string) UserPasswordVO {
	return UserPasswordVO{value: hash}
}

// Value nos permite obtener el hash para guardarlo en la DB
func (p UserPasswordVO) Value() string {
	return p.value
}
