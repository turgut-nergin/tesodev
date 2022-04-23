package models

import (
	"time"
)

type Customer struct {
	CustomerId string    `bson:"customerId"`
	Name       string    `bson:"name"`
	Email      string    `bson:"email"`
	Address    Address   `bson:"adress"`
	CreatedAdd time.Time `bson:",omitempty"`
	UpdatedAdd time.Time `bson:",omitempty"`
}
