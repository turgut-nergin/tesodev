package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
	"github.com/turgut-nergin/tesodev/repository/repo"
)

var GetCustomers = func(repository *repo.Repository) func(context *gin.Context) {
	return func(c *gin.Context) {
		req, err := repository.Get()
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
			customer = append(customer, response_models.Customer{
				Name:       cus.Name,
				CustomerId: cus.CustomerId,
				Email:      cus.Email,
				CreatedAdd: cus.CreatedAdd,
				UpdatedAdd: cus.UpdatedAdd,
				Address:    response_models.Address(cus.Address),
			})
		}

		c.JSON(http.StatusOK, customer)
	}
}
