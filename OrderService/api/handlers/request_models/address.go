package request_models

import validation "github.com/go-ozzo/ozzo-validation"

type Address struct {
	AddressLine string `json:"addressline"`
	City        string `json:"city"`
	Country     string `json:"country"`
	CityCode    int    `json:"citycode"`
}

func (address Address) validateAddress() error {
	return validation.ValidateStruct(&address,
		validation.Field(&address.AddressLine, validation.Required, validation.Length(5, 50)),
		validation.Field(&address.City, validation.Required, validation.Length(5, 50)),
		validation.Field(&address.Country, validation.Required, validation.Length(5, 50)),
		validation.Field(&address.CityCode, validation.Required, validation.Max(99999), validation.Min(10000)),
	)
}

func (address Address) Validate() error {
	return address.validateAddress()
}
