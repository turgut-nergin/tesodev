package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/api/handlers/lib"
	"github.com/turgut-nergin/tesodev/database"
)

var GetOrderById = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		orderId := c.Params.ByName("orderId")

		_, err := uuid.Parse(orderId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		req, err := r.GetOrderById(orderId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is not found"})
			return
		}

		order := lib.ResponseAssign(req)

		c.JSON(http.StatusOK, order)
	}
}
