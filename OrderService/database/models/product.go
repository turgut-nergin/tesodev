package models

type Product struct {
	ProductId string `bson:"productId"`
	ImageUrl  string `bson:"imageurl"`
	Name      string `bson:"name"`
}
