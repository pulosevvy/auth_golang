package validate

import "github.com/go-playground/validator"

type Validator struct {
	Valid *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.Valid.Struct(i)
}
