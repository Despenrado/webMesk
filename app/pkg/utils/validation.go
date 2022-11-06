package utils

import validation "github.com/go-ozzo/ozzo-validation"

func IsRequired(b bool) validation.RuleFunc {
	return func(val interface{}) error {
		if b {
			return validation.Validate(val, validation.Required)
		}
		return nil
	}
}
