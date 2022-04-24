package response_models

import (
	"time"
)

type Customer struct {
	CustomerId string    `bson:"customerId"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Address    Address   `json:"adress"`
	CreatedAdd time.Time `json:"createdAdd,omitempty"`
	UpdatedAdd time.Time `json:"updatedAdd,omitempty"`
}
