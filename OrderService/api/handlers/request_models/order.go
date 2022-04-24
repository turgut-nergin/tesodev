package request_models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Order struct {
	OrderId    string    `json:"orderid"`
	CustomerId string    `json:"customerid"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	Status     string    `json:"status"`
	Address    Address   `json:"address"`
	Product    Product   `json:"product"`
	CreatedAdd time.Time `json:",omitempty"`
	UpdatedAdd time.Time `json:",omitempty"`
}

func (order Order) validateOrder() error {
	return validation.ValidateStruct(&order,
		validation.Field(&order.Quantity, validation.Required, validation.Min(1)),
		validation.Field(&order.Price, validation.Required),
		validation.Field(&order.Status, validation.Required),
	)
}

func (order Order) Validate() error {
	err := order.validateOrder()
	if err != nil {
		return err
	}
	err = order.Address.Validate()
	if err != nil {
		return err
	}

	err = order.Product.Validate()
	if err != nil {
		return err
	}
	return nil
}
