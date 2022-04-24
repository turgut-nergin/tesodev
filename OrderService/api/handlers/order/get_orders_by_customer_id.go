package order

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/client"
	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
	"github.com/turgut-nergin/tesodev/database"
)

var GetOrdersByCustomerId = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		customerId := c.Params.ByName("customerId")
		err := c.ShouldBind(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		customerExist, statusError := client.Get(customerId)

		if statusError != nil {
			c.JSON(statusError.Code, gin.H{"error": statusError.Err.Error()})
			return
		}

		if !customerExist {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer ID is not registered"})
			return
		}

		req, err := r.GetOrdersByCustomerId(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is not found"})
			return
		}

		var orders []*response_models.Order
		for _, order := range req {
			orders = append(orders, &response_models.Order{
				OrderId:    order.OrderId,
				CustomerId: order.CustomerId,
				Quantity:   order.Quantity,
				Price:      order.Price,
				Status:     order.Status,
				Address:    response_models.Address(order.Address),
				Product:    response_models.Product(order.Product),
				CreatedAdd: time.Now(),
			})
		}

		// validCustomer := responseValidation.Customer{
		// 	Customer: *customer,
		// }

		// err = validCustomer.Validate()
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		c.JSON(http.StatusOK, orders)
	}
}
