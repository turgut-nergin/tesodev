package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
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

		customer := &response_models.Customer{
			Name:       req.Name,
			CustomerId: req.CustomerId,
			Email:      req.Email,
			CreatedAdd: req.CreatedAdd,
			UpdatedAdd: req.UpdatedAdd,
			Address:    response_models.Address(req.Address),
		}

		c.JSON(http.StatusOK, customer)
	}
}
