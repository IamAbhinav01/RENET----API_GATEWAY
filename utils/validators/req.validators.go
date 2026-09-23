package validators

import "github.com/go-playground/validator/v10"

var Validate validator.Validate

func NewValidate() validator.Validate{
	return *validator.New(
		validator.WithRequiredStructEnabled(),
	)
}

func Init(){
	Validate = NewValidate()
}