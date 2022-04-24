package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/database"
)

var DeleteOrderById = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		orderId := c.Param("orderId")

		if len(orderId) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id can not be empty!"})
			return
		}

		err := r.Delete(orderId)

		if err != nil {
			c.JSON(http.StatusBadRequest, false)
			return
		}

		c.JSON(http.StatusOK, true)
	}
}
