package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/lib/validations/customerValidation"
	"github.com/turgut-nergin/tesodev/repository/repo"
)

var ValidateCustomer = func(r *repo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		customerId := c.Params.ByName("customerId")

		err := c.ShouldBind(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		customer, err := r.GetByCustomerId(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, false)
			return
		}

		validCustomer := customerValidation.Customer{
			Customer: *customer,
		}

		err = validCustomer.Validate()

		//validate customer data or validate if  customer registry? I added both cases because I not  understand. I hope there isnot other case. :)
		if err != nil {
			c.JSON(http.StatusBadRequest, false)
		}

		c.JSON(http.StatusOK, true)
	}
}
