package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/api/client"
	"github.com/turgut-nergin/tesodev/api/handlers/request_models"
	"github.com/turgut-nergin/tesodev/database"
	"github.com/turgut-nergin/tesodev/database/models"
)

var CreateOrderHandler = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		customerId := c.Params.ByName("customerId")

		_, err := uuid.Parse(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		var req *request_models.Order
		err = c.ShouldBindJSON(&req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// orderValidations := []request_models.BaseRequest{req.Address, req.Product, req}

		// for _, a := range orderValidations {
		// 	a.Validate()
		// }

		err = req.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		customerExist, statusError := client.Get(customerId)

		if statusError != nil {
			c.JSON(statusError.Code, gin.H{"error": statusError.Err.Error()})
			return
		}

		if !customerExist {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer ID is not find"})
			return
		}

		order := &models.Order{
			CustomerId: customerId,
			Quantity:   req.Quantity,
			Price:      req.Price,
			Status:     req.Status,
			Address:    models.Address(req.Address),
			Product:    models.Product(req.Product),
		}

		orderR, err := r.Insert(order)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, orderR.OrderId)
	}
}
