package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
	"github.com/turgut-nergin/tesodev/database"
)

var GetOrdersHandler = func(database *database.Repository) func(context *gin.Context) {
	return func(c *gin.Context) {
		req, err := database.Get()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if req == nil {
			c.String(http.StatusNoContent, "Table is Empty")
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
				CreatedAdd: order.CreatedAdd,
				UpdatedAdd: order.UpdatedAdd,
			})
		}

		c.JSON(http.StatusOK, orders)
	}
}
