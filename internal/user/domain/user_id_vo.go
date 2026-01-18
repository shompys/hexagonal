package domain

import (
	"errors"
	"regexp"

	"github.com/shompys/hexagonal/pkg/validation"
)

// declaramos un value string y no un UserIDVO string para proteger el dato y esten obligados a usar constructor
type UserIDVO struct {
	value string
}

// constructor
func NewUserID(id string) (UserIDVO, error) {
	if err := validation.ValidateStringNotEmpty(ID, id); err != nil {
		return UserIDVO{}, err
	}
	if !isValidID(id) {
		return UserIDVO{}, errors.New("invalid ID format")
	}
	//retornamos copia porque un VO no tiene que poder ser cambiado desde otro lado
	return UserIDVO{
		value: id,
	}, nil
}

func isValidID(id string) bool {
	pattern := `^[\w\-]+$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(id)
}

func (v UserIDVO) Value() string {
	return v.value
}

func RestoreUserID(id string) UserIDVO {
	return UserIDVO{value: id} // No valida, solo asigna
}
