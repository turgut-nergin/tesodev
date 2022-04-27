package database

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/turgut-nergin/tesodev/database/lib"
	"github.com/turgut-nergin/tesodev/database/models"
	"github.com/turgut-nergin/tesodev/mongo"
	"gopkg.in/mgo.v2/bson"
)

var mockDB *Repository

func init() {
	dbModel := models.Repository{
		Name:           "testdb",
		CollectionName: "testcustomers",
	}

	client := mongo.NewClient("mongodb://localhost:27017")
	repo := New(client, dbModel)
	mockDB = repo
}

func dummyCustomerModel() *models.Customer {
	var customer models.Customer
	customer.Address.City = "Istanbul"
	customer.Name = "Nergin"
	customer.Email = "nergin.turgut@hotmail.com"
	customer.Address.Country = "Turkey"
	customer.Address.AddressLine = "yenişehir pendik"
	customer.Address.CityCode = 34188
	customer.CreatedAt = lib.TimeStampNow()
	return &customer
}

func cleanDB(db *mgo.Collection) error {
	_, err := db.RemoveAll(nil)
	return err
}

func getCustomer(query bson.M, db *mgo.Collection) (*models.Customer, error) {
	customer := models.Customer{}
	err := db.Find(query).One(&customer)
	return &customer, err

}

func sessionConfig() (*mgo.Collection, *mgo.Session) {
	mockSession := mockDB.mongoClient.NewSession()
	db := mockSession.DB("testdb").C("testcustomers")
	return db, mockSession
}

func insertDummyCustomer(db *mgo.Collection) (*models.Customer, error) {
	customer := dummyCustomerModel()
	err := db.Insert(customer)
	return customer, err
}

func insertDummyCustomers(customers []*models.Customer, db *mgo.Collection) {
	for _, customer := range customers {
		db.Insert(customer)
	}
}

func TestCreateCustomer(t *testing.T) {
	var dummyCustomer models.Customer
	dummyCustomer.Name = "nergissn"
	dummyCustomer.Address.City = "Istanbul"
	dummyCustomer.Address.Country = "Turkey"
	dummyCustomer.Address.AddressLine = "yenişehir pendik"
	dummyCustomer.Address.AddressLine = "34188"
	dummyCustomer.Email = "nergin.turgut@hotmail.com"

	reqCustomer, err := mockDB.Insert(&dummyCustomer)

	query := bson.M{"customerId": reqCustomer.CustomerId}

	if err != nil {
		t.Errorf("Insert dummy customer data error: %v\n", err)
		return
	}

	mockDB, mockSession := sessionConfig()
	defer mockSession.Close()

	got, err := getCustomer(query, mockDB)

	if err != nil {
		t.Errorf("Get customer error: %v\n", reqCustomer.CustomerId)
		return
	}

	if !reflect.DeepEqual(got, reqCustomer) {
		t.Errorf("got %v, want %v", got, reqCustomer)
		return
	}

	err = mockDB.Remove(query)

	if err != nil {
		t.Errorf(" Delete customer error: %v\n", err)
		return
	}

	t.Logf("Test create customer success")
}

func TestUpdateCustomer(t *testing.T) {
	mockCollection, mockSession := sessionConfig()
	defer mockSession.Close()

	dummyCustomer, err := insertDummyCustomer(mockCollection)

	if err != nil {
		t.Errorf("Create data error: %v\n", err)
		return
	}

	dummyCustomer.Email = "nergin.turgut@gmail.com"
	dummyCustomer.Name = "turgut"

	mockDB.Update(dummyCustomer.CustomerId, dummyCustomer)

	query := bson.M{"customerId": dummyCustomer.CustomerId}
	got, err := getCustomer(query, mockCollection)

	if err != nil {
		t.Errorf("Get customer error: %v\n", err)
		return
	}

	if !reflect.DeepEqual(got, dummyCustomer) {
		t.Errorf("got %v want %v", got, dummyCustomer)
		return
	}

	err = cleanDB(mockCollection)

	if err != nil {
		t.Errorf("Delete customer error: %v\n", err)
		return
	}
	t.Logf("Test update customer success :)")

}

func TestDeleteCustomer(t *testing.T) {
	mockCollection, mockSession := sessionConfig()
	defer mockSession.Close()

	dummyCustomer, err := insertDummyCustomer(mockCollection)

	if err != nil {
		t.Errorf("Insert data error: %v\n", err)
		return
	}

	err = mockDB.Delete(dummyCustomer.CustomerId)
	if err != nil {
		t.Errorf("Delete data error: %v\n", err)
	}

	query := bson.M{"customerId": dummyCustomer.CustomerId}
	_, err = getCustomer(query, mockCollection)

	if err == nil {
		t.Errorf("Test delete customer fail: %v\n", err)
	}

	cleanDB(mockCollection)

	t.Logf("Test delete customer success :)")

}

func TestIsValidCustomer(t *testing.T) {
	db, mockSession := sessionConfig()
	defer mockSession.Close()

	dummyCustomer, err := insertDummyCustomer(db)

	if err != nil {
		t.Errorf("Create data error: %v\n", err)
		return
	}

	got, err := mockDB.IdIsExist(dummyCustomer.CustomerId)

	if err != nil {
		t.Errorf("Get customer error: %v\n", err)
		return
	}

	if got != true {
		t.Errorf("got %v want %v", got, true)
		return
	}

	cleanDB(db)

	t.Logf("Test is valid customer id success")

}

func TestGetCustomerById(t *testing.T) {
	db, mockSession := sessionConfig()
	defer mockSession.Close()

	dummyCustomer, err := insertDummyCustomer(db)

	if err != nil {
		t.Errorf("Create data error: %v\n", err)
		return
	}

	got, err := mockDB.GetByCustomerId(dummyCustomer.CustomerId)

	if err != nil {
		t.Errorf("Get customer error: %v\n", err)
		return
	}

	if !reflect.DeepEqual(got, dummyCustomer) {
		t.Errorf("got %v want %v", got, dummyCustomer)
		return
	}

	cleanDB(db)

	t.Logf("Test get customer by id success")
}

func TestGetCustomers(t *testing.T) {
	db, mockSession := sessionConfig()
	defer mockSession.Close()
	dummyData := dummyCustomerModel()
	dummyCustomers := []*models.Customer{dummyData, dummyData}
	insertDummyCustomers(dummyCustomers, db)

	got, err := mockDB.Get()

	if err != nil {
		t.Errorf("Get customer error: %v\n", err)
		return
	}

	if !reflect.DeepEqual(got, dummyCustomers) {
		if len(got) == len(dummyCustomers) {
			for i := range dummyCustomers {
				t.Errorf("got %v want %v", got[i], dummyCustomers[i])
			}
			return
		}
		t.Errorf("Failed to get customers data")
		return
	}

	cleanDB(db)

	t.Logf("Test get customers success")
}
