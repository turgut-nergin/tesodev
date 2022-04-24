package database

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/turgut-nergin/tesodev/database/models"
	"github.com/turgut-nergin/tesodev/mongo"
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

func (r *Repository) IdIsExist(id string) (bool, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()
	query := bson.M{"customerId": id}

	count, err := session.
		DB("tesodev").
		C("customer").Find(query).Count()

	if err != nil {
		return false, err
	}
	return count > 0, nil
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
