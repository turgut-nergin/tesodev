package request_models

import (
	"time"
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
