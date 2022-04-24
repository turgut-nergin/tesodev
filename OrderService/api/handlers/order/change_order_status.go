package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/database"
)

var ChangeOrderStatus = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		orderId := c.Params.ByName("orderId")

		_, err := uuid.Parse(orderId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		status := c.Query("status")
		if len(status) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "status can not be empty"})
			return
		}

		changed, err := r.UpdateOrderStatus(orderId, status)

		if err != nil {
			c.JSON(http.StatusBadRequest, false)
			return
		}

		c.JSON(http.StatusOK, changed)
	}
}
