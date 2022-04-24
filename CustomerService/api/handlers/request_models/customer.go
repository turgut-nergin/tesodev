package request_models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Customer struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Address    Address   `json:"address"`
	CreatedAdd time.Time `json:",omitempty"`
	UpdatedAdd time.Time `json:",omitempty"`
}

func (customer Customer) validateCustomer() error {
	return validation.ValidateStruct(&customer,
		validation.Field(&customer.Email, validation.Required, is.Email),
		validation.Field(&customer.Name, validation.Required, validation.Length(1, 25)),
	)
}

func (customer Customer) Validate() error {

	err := customer.validateCustomer()
	if err != nil {
		return err
	}

	err = customer.Address.Validate()

	if err != nil {
		return err
	}
	return nil

}
