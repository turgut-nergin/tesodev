package models

type Order struct {
	OrderId    string  `bson:"orderId,omitempty"`
	CustomerId string  `bson:"customerId,omitempty"`
	Quantity   int     `bson:"quantity"`
	Price      float64 `bson:"price"`
	Status     string  `bson:"status"`
	Address    Address `bson:"address"`
	Product    Product `bson:"product"`
	CreatedAt  int64   `bson:",omitempty"`
	UpdatedAt  int64   `bson:",omitempty"`
}
