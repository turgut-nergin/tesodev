package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
	"github.com/turgut-nergin/tesodev/database"
)

var GetOrderById = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		orderId := c.Params.ByName("orderId")
		err := c.ShouldBind(orderId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		req, err := r.GetOrderById(orderId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is not found"})
			return
		}

		order := &response_models.Order{
			OrderId:    req.OrderId,
			CustomerId: req.CustomerId,
			Quantity:   req.Quantity,
			Price:      req.Price,
			Status:     req.Status,
			Address:    response_models.Address(req.Address),
			Product:    response_models.Product(req.Product),

			CreatedAdd: req.CreatedAdd,
			UpdatedAdd: req.UpdatedAdd,
		}

		// validCustomer := responseValidation.Customer{
		// 	Customer: *customer,
		// }

		// err = validCustomer.Validate()
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		c.JSON(http.StatusOK, order)
	}
}
