package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/api/handlers/lib"
	"github.com/turgut-nergin/tesodev/database"
)

var GetCustomerByCustomerId = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		customerId := c.Params.ByName("customerId")

		_, err := uuid.Parse(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		req, err := r.GetByCustomerId(customerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is not found"})
			return
		}

		if req == nil {
			c.String(http.StatusNotFound, "Table is Empty")
			return
		}

		customer := lib.ResponseAssign(req)

		c.JSON(http.StatusOK, customer)
	}
}
