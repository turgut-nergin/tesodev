package customer

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/api/handlers/request_models"
	"github.com/turgut-nergin/tesodev/api/lib/validations/customerValidation"
	"github.com/turgut-nergin/tesodev/api/lib/validations/requestValidation"
	"github.com/turgut-nergin/tesodev/repository/models"
	"github.com/turgut-nergin/tesodev/repository/repo"
)

var CreateCustomer = func(r *repo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		var req *request_models.Customer
		err := c.ShouldBindJSON(&req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not bind the request!"})
			return
		}

		validRequest := requestValidation.Customer{
			Customer: *req,
		}

		err = validRequest.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//#TODO: you must converted uuid to bson type :)
		customer := &models.Customer{
			Name:       req.Name,
			CustomerId: uuid.New().String(),
			Email:      req.Email,
			Address:    models.Address(req.Address),
			CreatedAdd: time.Now(),
		}

		validCustomer := customerValidation.Customer{
			Customer: *customer,
		}

		err = validCustomer.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		customerR, err := r.Insert(customer)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusAccepted, customerR.CustomerId)
	}
}
