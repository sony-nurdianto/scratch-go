package model

import "github.com/sony-nurdianto/scratch-go/graph/validator"

func (r RegisterUser) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", r.Email)
	v.IsEmail("email", r.Email)

	v.Required("password", r.Password)
	v.MinLength("password", r.Password, 6)

	v.Required("confirmPassword", r.ConfirmPassword)
	v.EqualToField("confirmPassword", r.ConfirmPassword, "password", r.Password)

	v.Required("user_name", r.UserName)
	v.MinLength("user_name", r.UserName, 2)

	v.Required("first_name", r.FirstName)
	v.MinLength("first_name", r.FirstName, 2)

	v.Required("last_name", r.LastName)
	v.MinLength("last_name", r.LastName, 2)

	return v.IsValid(), v.Errors
}

func (l LoginUser) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", l.Email)
	v.IsEmail("email", l.Email)

	v.Required("password", l.Password)

	return v.IsValid(), v.Errors
}
