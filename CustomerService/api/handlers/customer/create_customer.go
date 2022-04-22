package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/api/handlers/request_models"
	"github.com/turgut-nergin/tesodev/repository/models"
	repositorty "github.com/turgut-nergin/tesodev/repository/repo"
)

var CreateCustomer = func(r *repositorty.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		var req *request_models.Customer
		err := c.ShouldBindJSON(&req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not bind the request!"})
			return
		}

		//#TODO: you must converted uuid to bson type :)
		customer := &models.Customer{
			Name:    req.Name,
			UserID:  uuid.New().String(),
			Email:   req.Email,
			Address: models.Address(req.Address),
		}

		customerR, err := r.Insert(customer)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusAccepted, customerR)
	}
}
