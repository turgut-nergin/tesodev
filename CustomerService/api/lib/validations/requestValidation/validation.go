package requestValidation

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/turgut-nergin/tesodev/api/handlers/request_models"
)

type Customer struct {
	Customer request_models.Customer
}

func ValidateAddress(validate request_models.Address) error {
	return validation.ValidateStruct(&validate,
		validation.Field(&validate.AddressLine, validation.Required, validation.Length(5, 50)),
		validation.Field(&validate.City, validation.Required, validation.Length(5, 50)),
		validation.Field(&validate.Country, validation.Required, validation.Length(5, 50)),
		validation.Field(&validate.CityCode, validation.Required, validation.Max(99999), validation.Min(10000)),
	)
}

func ValidateCustomer(validate request_models.Customer) error {
	return validation.ValidateStruct(&validate,
		validation.Field(&validate.Email, validation.Required, is.Email),
		validation.Field(&validate.Name, validation.Required, validation.Length(5, 25)),
	)
}

func (customer Customer) Validate() error {
	err := ValidateCustomer(customer.Customer)
	if err != nil {
		return err
	}
	err = ValidateAddress(customer.Customer.Address)
	if err != nil {
		return err
	}
	return nil

}
