package repo

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/turgut-nergin/tesodev/mongo"
	"github.com/turgut-nergin/tesodev/repository/models"
)

type Repository struct {
	mongoClient *mongo.Client
}

func (r *Repository) GetByCustomerId(id string) (*models.Customer, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()
	query := bson.M{"customerId": id}
	var customer *models.Customer
	err := session.
		DB("tesodev").
		C("customer").
		Find(query).
		One(&customer)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *Repository) Get() ([]models.Customer, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()
	var customer []models.Customer
	err := session.
		DB("tesodev").
		C("customer").
		Find(nil).
		Limit(100).
		Iter().
		All(&customer)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *Repository) Insert(customer *models.Customer) (*models.Customer, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()

	err := session.
		DB("tesodev").
		C("customer").
		Insert(customer)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *Repository) Update(id string, customer *models.Customer) (*models.Customer, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()
	selector := bson.M{"customerId": id}

	// updateCustomer := bson.D{
	// 	{"$set", bson.D{
	// 		{"name", customer.Name},
	// 		{"customerId", customer.CustomerId},
	// 		{"email", customer.Email},
	// 		{"address", bson.D{
	// 			{"addressline", customer.Address.AddressLine},
	// 			{"city", customer.Address.City},
	// 			{"country", customer.Address.Country},
	// 			{"citycode", customer.Address.CityCode},
	// 		}},
	// 		{"updatedAdd", time.Now()},
	// 	}}}

	// updateCustomer := bson.M{
	// 	"$set": bson.M{
	// 		"name":       customer.Name,
	// 		"customerId": customer.CustomerId,
	// 		"email":      customer.Email,
	// 		"address": bson.M{
	// 			"addressline": customer.Address.AddressLine,
	// 			"city":        customer.Address.City,
	// 			"country":     customer.Address.Country,
	// 			"citycode":    customer.Address.CityCode,
	// 		},
	// 		"updatedAdd": time.Now(),
	// 	}}

	customer.UpdatedAdd = time.Now()

	err := session.
		DB("tesodev").
		C("customer").Update(selector, bson.M{
		"$set": customer})

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *Repository) Delete(id string) error {
	var session = r.mongoClient.NewSession()
	query := bson.M{"customerId": id}

	defer session.Close()

	err := session.
		DB("tesodev").
		C("customer").
		Remove(query)

	if err != nil {
		return err
	}
	return nil
}

func New(mongoClient *mongo.Client) *Repository {
	repo := Repository{mongoClient}
	return &repo
}
