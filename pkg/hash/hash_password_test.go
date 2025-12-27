package hash

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "securepassword123"

	hasher := PasswordHasher{}
	// Generar el hash
	hash, err := hasher.HashPassword(password)

	if err != nil {
		t.Fatalf("Error al generar el hash: %v", err)
	}

	// Verificar que el hash no sea igual a la contraseña original
	if hash == password {
		t.Errorf("El hash no debería ser igual a la contraseña original")
	}

	// Validar el hash generado con bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("El hash generado no valida correctamente la contraseña: %v", err)
	}
}
