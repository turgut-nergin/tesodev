package response_models

import (
	"time"
)

type Customer struct {
	CustomerId string    `bson:"customerId"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Address    Address   `json:"adress"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}
