package database

import (
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/database/lib"
	"github.com/turgut-nergin/tesodev/database/models"
	"github.com/turgut-nergin/tesodev/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct {
	mongoClient *mongo.Client
	DB          models.Repository
}

func (r *Repository) GetOrdersByCustomerId(id string) ([]*models.Order, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()
	var order []*models.Order
	query := bson.M{"customerId": id}

	err := session.
		DB(r.DB.Name).
		C(r.DB.CollectionName).
		Find(query).
		Limit(100).
		Iter().
		All(&order)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *Repository) Get() ([]*models.Order, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()
	var order []*models.Order
	err := session.
		DB(r.DB.Name).
		C(r.DB.CollectionName).
		Find(nil).
		Limit(100).
		Iter().
		All(&order)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *Repository) Insert(order *models.Order) (*models.Order, error) {
	var session = r.mongoClient.NewSession()

	defer session.Close()

	order.OrderId = uuid.NewString()
	order.CreatedAt = lib.TimeStampNow()

	err := session.
		DB(r.DB.Name).
		C(r.DB.CollectionName).
		Insert(order)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *Repository) UpdateOrderStatus(id string, status string) (bool, error) {
	var session = r.mongoClient.NewSession()

	defer session.Close()

	selector := bson.M{"orderId": id}

	err := session.
		DB(r.DB.Name).
		C(r.DB.CollectionName).Update(selector, bson.M{
		"$set": bson.M{"status": status, "updatedat": lib.TimeStampNow()}})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Repository) Update(id string, order *models.Order) (bool, error) {
	var session = r.mongoClient.NewSession()

	defer session.Close()

	selector := bson.M{"orderId": id}
	order.UpdatedAt = lib.TimeStampNow()

	err := session.
		DB(r.DB.Name).
		C(r.DB.CollectionName).Update(selector, bson.M{
		"$set": order})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Repository) GetOrderById(id string) (*models.Order, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()

	query := bson.M{"orderId": id}

	var order *models.Order

	err := session.
		DB(r.DB.Name).
		C(r.DB.CollectionName).
		Find(query).
		One(&order)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *Repository) Delete(id string) error {
	var session = r.mongoClient.NewSession()
	query := bson.M{"orderId": id}

	defer session.Close()

	err := session.
		DB(r.DB.Name).
		C(r.DB.CollectionName).
		Remove(query)

	if err != nil {
		return err
	}
	return nil
}

func New(mongoClient *mongo.Client, db models.Repository) *Repository {
	repo := Repository{
		mongoClient: mongoClient,
		DB:          db,
	}
	return &repo
}
