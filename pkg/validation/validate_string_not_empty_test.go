package validation

import (
	"fmt"
	"testing"
)

func TestValidateStringNotEmpty(t *testing.T) {
	tests := []struct {
		fieldName string
		value     string
		expected  error
	}{
		{"Name", "John", nil},                            // Caso válido
		{"Name", "", fmt.Errorf("Name cannot be empty")}, // Caso inválido
	}

	for _, test := range tests {
		err := ValidateStringNotEmpty(test.fieldName, test.value)
		if err != nil && err.Error() != test.expected.Error() {
			t.Errorf("validateStringNotEmpty(%q, %q) = %v; want %v", test.fieldName, test.value, err, test.expected)
		}
		if err == nil && test.expected != nil {
			t.Errorf("validateStringNotEmpty(%q, %q) = nil; want %v", test.fieldName, test.value, test.expected)
		}
	}
}
