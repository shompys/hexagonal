package validation

import "testing"

func TestIsEmailValid(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},           // Caso válido
		{"user+alias@gmail.com", true},       // Caso válido con '+'
		{"invalid-email", false},             // Sin '@' ni dominio
		{"@example.com", false},              // Sin parte local
		{"user@.com", false},                 // Dominio inválido
		{"user@domain", false},               // Sin extensión de dominio
		{"user@domain.c", false},             // Extensión de dominio muy corta
		{"user@domain.com", true},            // Caso válido
		{"user@sub.domain.com", true},        // Subdominio válido
		{"pirulo@dominio..com", false},       // Dominio con puntos consecutivos
		{"pirulo@dominio.com.com.com", true}, // Dominio con múltiples subdominios
		{"pirulo@dominio.co", true},          // Dominio con extensión de dos letras
	}

	for _, test := range tests {
		result := IsEmailValid(test.email)
		if result != test.expected {
			t.Errorf("IsEmailValid(%q) = %v; want %v", test.email, result, test.expected)
		}
	}
}
