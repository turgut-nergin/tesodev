package test

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/order"
	"github.com/turgut-nergin/tesodev/database"
)

type ContextMock struct {
	JSONCalled bool
}

func (c *ContextMock) JSON(code int, obj interface{}) {
	c.JSONCalled = true
}

type JSONer interface {
	JSON(code int, obj interface{})
}

func CreateOrder(c JSONer) {
	var repository *database.Repository
	c.JSON(200, order.GetOrdersHandler(repository))
}

func TestControllerGetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c := &ContextMock{false}

	CreateOrder(c)
	fmt.Println(c)

	if !c.JSONCalled {
		fmt.Print("Failed")
	} else {
		fmt.Print("Tested successfully")
	}
}
