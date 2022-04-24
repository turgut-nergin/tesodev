package response_models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	ID         bson.ObjectId `json:"id,omitempty"`
	OrderId    string        `json:"orderid"`
	CustomerId string        `json:"customerid"`
	Quantity   int           `json:"quantity"`
	Price      float64       `json:"price"`
	Status     string        `json:"status"`
	Address    Address       `json:"address"`
	Product    Product       `json:"product"`
	CreatedAdd time.Time     `json:"createdAdd,omitempty"`
	UpdatedAdd time.Time     `json:"updatedAdd,omitempty"`
}
