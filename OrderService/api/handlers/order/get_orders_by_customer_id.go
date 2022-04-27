package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/api/client"
	"github.com/turgut-nergin/tesodev/api/handlers/lib"
	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
	"github.com/turgut-nergin/tesodev/database"
)

var GetOrdersByCustomerId = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		customerId := c.Params.ByName("customerId")

		_, err := uuid.Parse(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		customerExist, statusError := client.Get(customerId)

		if statusError != nil {
			c.JSON(statusError.Code, gin.H{"error": statusError.Err.Error()})
			return
		}

		if !customerExist {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer ID is not found"})
			return
		}

		req, err := r.GetOrdersByCustomerId(customerId)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{"error": "Table is empty"})
			return
		}

		var orders []*response_models.Order
		for _, order := range req {
			reponse_model := lib.ResponseAssign(order)
			orders = append(orders, reponse_model)
		}

		c.JSON(http.StatusOK, orders)
	}
}
