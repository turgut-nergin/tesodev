package models

type Customer struct {
	CustomerId string  `bson:"customerId"`
	Name       string  `bson:"name"`
	Email      string  `bson:"email"`
	Address    Address `bson:"adress"`
	CreatedAt  int64   `bson:",omitempty"`
	UpdatedAt  int64   `bson:",omitempty"`
}
