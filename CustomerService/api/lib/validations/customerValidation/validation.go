package customerValidation

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/turgut-nergin/tesodev/database/models"
)

type Customer struct {
	Customer models.Customer
}

func ValidateAddress(validate models.Address) error {
	return validation.ValidateStruct(&validate,
		validation.Field(&validate.AddressLine, validation.Required, validation.Length(5, 50)),
		validation.Field(&validate.City, validation.Required, validation.Length(5, 50)),
		validation.Field(&validate.Country, validation.Required, validation.Length(5, 50)),
		validation.Field(&validate.CityCode, validation.Required, validation.Max(99999), validation.Min(10000)),
	)
}

func ValidateCustomer(validate models.Customer) error {
	return validation.ValidateStruct(&validate,
		validation.Field(&validate.CustomerId, validation.Required, is.UUIDv4),
		validation.Field(&validate.Email, validation.Required, is.Email),
		validation.Field(&validate.Name, validation.Required, validation.Length(5, 25)),
		//#TODO: It will be edit !!.
		validation.Field(&validate.CreatedAdd, validation.Required),
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
