package validation

import "fmt"

func ValidateStringNotEmpty(fieldName, value string) error {
	if value == "" {
		return fmt.Errorf("%s cannot be empty", fieldName)
	}
	return nil
}
