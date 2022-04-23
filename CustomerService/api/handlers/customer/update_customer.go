package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/request_models"
	"github.com/turgut-nergin/tesodev/api/lib/validations/requestValidation"
	"github.com/turgut-nergin/tesodev/repository/models"
	"github.com/turgut-nergin/tesodev/repository/repo"
)

var UpdateCustomerHandler = func(r *repo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		customerId := c.Param("customerId")

		if customerId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user ID can not be empty"})
			return
		}

		var req request_models.Customer
		err := c.ShouldBindJSON(&req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		validRequest := requestValidation.Customer{
			Customer: req,
		}

		err = validRequest.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		address := &models.Address{
			AddressLine: req.Address.AddressLine,
			City:        req.Address.City,
			CityCode:    req.Address.CityCode,
			Country:     req.Address.Country,
		}
		customer := &models.Customer{
			Name:       req.Name,
			CustomerId: customerId,
			Email:      req.Email,
			Address:    *address,
		}

		_, err = r.Update(customerId, customer)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusAccepted, true)
	}
}
