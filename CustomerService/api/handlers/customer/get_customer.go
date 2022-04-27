package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/turgut-nergin/tesodev/api/handlers/lib"
	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
	"github.com/turgut-nergin/tesodev/database"
)

var GetCustomers = func(database *database.Repository) func(context *gin.Context) {
	return func(c *gin.Context) {
		req, err := database.Get()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if req == nil {
			c.String(http.StatusNotFound, "Table is Empty")
			return
		}

		var customer []response_models.Customer
		for _, cus := range req {
			response := lib.ResponseAssign(cus)
			customer = append(customer, *response)
		}

		c.JSON(http.StatusOK, customer)
	}
}
