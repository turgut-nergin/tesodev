package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/database"
)

var DeleteCustomerById = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		customerId := c.Param("customerId")

		_, err := uuid.Parse(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = r.Delete(customerId)

		if err != nil {
			c.JSON(http.StatusNoContent, false)
			return
		}

		c.JSON(http.StatusOK, true)
	}
}
