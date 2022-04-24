package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/database"
)

var ValidateCustomer = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		customerId := c.Param("customerId")

		if customerId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Customer Id can not be empty!"})
			return
		}

		isExist, err := r.IdIsExist(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusOK, isExist)
	}
}
