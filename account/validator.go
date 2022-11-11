package account

import (
	"github.com/go-playground/validator"
)

func validateCurrency(fl validator.FieldLevel) bool {
	currency := fl.Field().Interface().(string)

	if currency != "AZN" && currency != "USD" && currency != "EUR" {
		return false
	}
	return true
}
