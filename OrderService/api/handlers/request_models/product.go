package request_models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Product struct {
	ProductId string `json:"ProductId"`
	ImageUrl  string `json:"imageurl"`
	Name      string `json:"name"`
}

func (product Product) validateProduct() error {
	return validation.ValidateStruct(&product,
		validation.Field(&product.ProductId, validation.Required, is.UUIDv4),
		validation.Field(&product.Name, validation.Required),
		validation.Field(&product.ImageUrl, validation.Required, is.URL),
	)
}

func (product Product) Validate() error {
	return product.validateProduct()
}
