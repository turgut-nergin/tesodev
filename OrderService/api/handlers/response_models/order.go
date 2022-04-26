package response_models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	ID         bson.ObjectId `json:"_id,omitempty"`
	OrderId    string        `json:"orderid"`
	CustomerId string        `json:"customerid"`
	Quantity   int           `json:"quantity"`
	Price      float64       `json:"price"`
	Status     string        `json:"status"`
	Address    Address       `json:"address"`
	Product    Product       `json:"product"`
	CreatedAt  time.Time     `json:"createdAt,omitempty"`
	UpdatedAt  time.Time     `json:"updatedAt,omitempty"`
}
