package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/lib"
	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
	"github.com/turgut-nergin/tesodev/database"
)

var GetOrdersHandler = func(r *database.Repository) func(context *gin.Context) {
	return func(c *gin.Context) {
		req, err := r.Get()
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
			reponse_model := lib.ResponseAssign(order)
			orders = append(orders, reponse_model)
		}

		c.JSON(http.StatusOK, orders)
	}
}
