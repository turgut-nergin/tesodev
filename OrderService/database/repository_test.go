package database

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/database/lib"
	"github.com/turgut-nergin/tesodev/database/models"
	"github.com/turgut-nergin/tesodev/mongo"
	"gopkg.in/mgo.v2/bson"
)

var mockDB *Repository

func init() {
	dbModel := models.Repository{
		Name:           "testdb",
		CollectionName: "testorders",
	}

	client := mongo.NewClient("mongodb://localhost:27017")
	repo := New(client, dbModel)
	mockDB = repo
}

func dummyOrderModel() *models.Order {
	var order models.Order
	order.OrderId = "26040276-b961-19a9-61c3-19280a16c611"
	order.Address.City = "Istanbul"
	order.Address.Country = "Turkey"
	order.Address.AddressLine = "yenişehir pendik"
	order.Address.CityCode = 34188
	order.CustomerId = "1704027f-a962-49f9-5c74-79180ae6a918"
	order.Product.ProductId = "1918011e-49a2-6qfe-6c10-11189ae5a9a9"
	order.Product.ImageUrl = "https://cdn.dsmcdn.com/mnresize/1200/1800/ty106/product/media/images/20210421/6/81798177/165468269/1/1_org_zoom.jpg"
	order.Product.Name = "Harry Potter Diagon Alley Book"
	order.Status = "order being prepared"
	order.CreatedAt = lib.TimeStampNow()
	return &order
}

func cleanDB(db *mgo.Collection) error {
	_, err := db.RemoveAll(nil)
	return err
}

func getOrder(query bson.M, db *mgo.Collection) (*models.Order, error) {
	order := models.Order{}
	err := db.Find(query).One(&order)
	return &order, err

}

func sessionConfig() (*mgo.Collection, *mgo.Session) {
	mockSession := mockDB.mongoClient.NewSession()
	db := mockSession.DB("testdb").C("testorders")
	return db, mockSession
}

func insertDummyOrder(db *mgo.Collection) (*models.Order, error) {
	order := dummyOrderModel()
	err := db.Insert(order)
	return order, err
}

func insertDummyOrders(orders []*models.Order, db *mgo.Collection) {
	for _, order := range orders {
		db.Insert(order)
	}
}

func TestCreateOrder(t *testing.T) {
	var order models.Order
	order.Address.City = "Istanbul"
	order.Address.Country = "Turkey"
	order.Address.AddressLine = "yenişehir pendik"
	order.Address.AddressLine = "34188"
	order.CustomerId = uuid.New().String()
	order.Product.ProductId = uuid.New().String()
	order.Product.ImageUrl = "https://cdn.dsmcdn.com/mnresize/1200/1800/ty106/product/media/images/20210421/6/81798177/165468269/1/1_org_zoom.jpg"
	order.Product.Name = "Harry Potter Diagon Alley Book"
	order.Status = "sipariş yolda"

	reqOrder, err := mockDB.Insert(&order)

	query := bson.M{"orderId": reqOrder.OrderId}

	if err != nil {
		t.Errorf("Insert dummy order data error: %v\n", err)
		return
	}

	mockDB, mockSession := sessionConfig()
	defer mockSession.Close()

	got := models.Order{}
	err = mockDB.Find(bson.M{"orderId": reqOrder.OrderId}).One(&got)

	if err != nil {
		t.Errorf("TestCreateOrder: Get order error: %v\n", order.OrderId)
		return
	}

	if !reflect.DeepEqual(got, order) {
		t.Errorf("got %v, want %v", got.OrderId, order.CustomerId)
		return
	}

	err = mockDB.Remove(query)

	if err != nil {
		t.Errorf("TestCreateOrder: Delete order error: %v\n", err)
		return
	}

	t.Logf("Success")
}

func TestUpdateOrder(t *testing.T) {
	mockCollection, mockSession := sessionConfig()
	defer mockSession.Close()

	order, err := insertDummyOrder(mockCollection)

	if err != nil {
		t.Errorf("Create data error: %v\n", err)
		return
	}

	order.UpdatedAt = lib.TimeStampNow()
	order.Price = 29.90
	order.Quantity = 2
	mockDB.Update(order.OrderId, order)

	query := bson.M{"orderId": order.OrderId}
	got, err := getOrder(query, mockCollection)

	if err != nil {
		t.Errorf("Get order error: %v\n", err)
		return
	}

	if !reflect.DeepEqual(got, order) {
		t.Errorf("got %v want %v", got, order)
		return
	}

	err = cleanDB(mockCollection)

	if err != nil {
		t.Errorf("Delete order error: %v\n", err)
		return
	}
	t.Logf("Test Update Order Success :)")

}

func TestDeleteOrder(t *testing.T) {
	mockCollection, mockSession := sessionConfig()
	defer mockSession.Close()

	dummyOrder, err := insertDummyOrder(mockCollection)

	if err != nil {
		t.Errorf("Delete data error: %v\n", err)
		return
	}

	err = mockDB.Delete(dummyOrder.OrderId)
	if err != nil {
		t.Errorf("Delete data error: %v\n", err)
	}

	query := bson.M{"orderId": dummyOrder.OrderId}
	_, err = getOrder(query, mockCollection)

	if err == nil {
		t.Errorf("\tTest Delete Order fail: %v\n", err)
	}

	cleanDB(mockCollection)

	t.Logf("Test Delete Order Success")

}

func TestChangeOrderStatus(t *testing.T) {
	db, mockSession := sessionConfig()
	defer mockSession.Close()

	order, err := insertDummyOrder(db)

	if err != nil {
		t.Errorf("Create data error: %v\n", err)
		return
	}

	order.UpdatedAt = lib.TimeStampNow()
	order.Status = "Teslim edildi"
	mockDB.Update(order.OrderId, order)

	query := bson.M{"orderId": order.OrderId}
	got, err := getOrder(query, db)

	if err != nil {
		t.Errorf("Get order error: %v\n", err)
		return
	}

	if !reflect.DeepEqual(got, order) {
		t.Errorf("got %v want %v", got, order)
		return
	}

	cleanDB(db)

	t.Logf("Test Change Order Status Success")

}

func TestGetOrderById(t *testing.T) {
	db, mockSession := sessionConfig()
	defer mockSession.Close()

	order, err := insertDummyOrder(db)

	if err != nil {
		t.Errorf("Create data error: %v\n", err)
		return
	}

	got, err := mockDB.GetOrderById(order.OrderId)

	if err != nil {
		t.Errorf("Get order error: %v\n", err)
		return
	}

	if !reflect.DeepEqual(got, order) {
		t.Errorf("got %v want %v", got, order)
		return
	}

	cleanDB(db)

	t.Logf("Test Get Order By ID Success :)")
}

func TestGetOrdersByCustomerID(t *testing.T) {
	db, mockSession := sessionConfig()
	defer mockSession.Close()
	dummyData := dummyOrderModel()
	orders := []*models.Order{dummyData, dummyData}
	insertDummyOrders(orders, db)

	got, err := mockDB.GetOrdersByCustomerId(dummyData.CustomerId)

	if err != nil {
		t.Errorf("Get order error: %v\n", err)
		return
	}

	if !reflect.DeepEqual(got, orders) {
		if len(got) == len(orders) {
			for i := range orders {
				t.Errorf("got %v want %v", got[i], orders[i])
			}
			return
		}
		t.Errorf("Failed to get orders data")
		return
	}

	cleanDB(db)

	t.Logf("Test Get Orders by Customer ID Success")
}

func TestGetOrders(t *testing.T) {
	db, mockSession := sessionConfig()
	defer mockSession.Close()
	dummyData := dummyOrderModel()
	orders := []*models.Order{dummyData, dummyData}
	insertDummyOrders(orders, db)

	got, err := mockDB.Get()

	if err != nil {
		t.Errorf("Get order error: %v\n", err)
		return
	}

	if !reflect.DeepEqual(got, orders) {
		if len(got) == len(orders) {
			for i := range orders {
				t.Errorf("got %v want %v", got[i], orders[i])
			}
			return
		}
		t.Errorf("Failed to get orders data")
		return
	}

	cleanDB(db)

	t.Logf("Test Get Orders by Customer ID Success")
}
