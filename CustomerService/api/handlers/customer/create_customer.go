package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/request_models"
	"github.com/turgut-nergin/tesodev/database"
	"github.com/turgut-nergin/tesodev/database/models"
)

var CreateCustomer = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		var req *request_models.Customer
		err := c.ShouldBindJSON(&req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not bind the request!"})
			return
		}

		err = req.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//#TODO: you must converted uuid to bson type :)
		customer := &models.Customer{
			Name:    req.Name,
			Email:   req.Email,
			Address: models.Address(req.Address),
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		customerR, err := r.Insert(customer)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, customerR.CustomerId)
	}
}
