package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/database"
)

var DeleteOrderById = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		orderId := c.Param("orderId")

		_, err := uuid.Parse(orderId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = r.Delete(orderId)

		if err != nil {
			c.JSON(http.StatusBadRequest, false)
			return
		}

		c.JSON(http.StatusOK, true)
	}
}
