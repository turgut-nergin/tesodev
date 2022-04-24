package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/database"
)

var ChangeOrderStatus = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		orderId := c.Params.ByName("orderId")

		err := c.ShouldBind(orderId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
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
