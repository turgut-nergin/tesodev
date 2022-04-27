package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/api/handlers/request_models"
	"github.com/turgut-nergin/tesodev/database"
	"github.com/turgut-nergin/tesodev/database/models"
)

var UpdateCustomerHandler = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		customerId := c.Param("customerId")

		_, err := uuid.Parse(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		var req request_models.Customer
		err = c.ShouldBindJSON(&req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = req.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		customer := &models.Customer{
			Name:       req.Name,
			CustomerId: customerId,
			Email:      req.Email,
			Address:    models.Address(req.Address),
		}

		_, err = r.Update(customerId, customer)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, true)
	}
}
