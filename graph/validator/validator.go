package validator

//Validator Struct
type Validator struct {
	Errors map[string]string
}

//Validation interface
type Validation interface {
	Validate() (bool, map[string]string)
}

//New To declare New Validator
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

//IsValid to check is Valid
func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}
