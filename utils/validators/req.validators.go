package validators

import "github.com/go-playground/validator/v10"

var Validate = NewValidate()

func NewValidate() validator.Validate{
	return *validator.New(
		validator.WithRequiredStructEnabled(),
	)
}
