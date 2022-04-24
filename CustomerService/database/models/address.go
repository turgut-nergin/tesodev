package models

type Address struct {
	AddressLine string `bson:"addressline"`
	City        string `bson:"city"`
	Country     string `bson:"country"`
	CityCode    int    `bson:"citycode"`
}
