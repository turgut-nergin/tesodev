package models

import (
	"time"
)

type Order struct {
	OrderId    string    `bson:"orderId,omitempty"`
	CustomerId string    `bson:"customerId,omitempty"`
	Quantity   int       `bson:"quantity"`
	Price      float64   `bson:"price"`
	Status     string    `bson:"status"`
	Address    Address   `bson:"address"`
	Product    Product   `bson:"product"`
	CreatedAdd time.Time `bson:",omitempty"`
	UpdatedAdd time.Time `bson:",omitempty"`
}
