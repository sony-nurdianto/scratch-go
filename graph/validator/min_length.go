package validator

import "fmt"

//MinLength is to validate minimum length
func (v *Validator) MinLength(field, value string, minCharacter int) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if len(value) < minCharacter {
		v.Errors[field] = fmt.Sprintf("%s must be at least(%d) characters long", field, minCharacter)

		return false
	}

	return true
}
