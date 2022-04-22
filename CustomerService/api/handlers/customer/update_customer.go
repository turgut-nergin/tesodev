package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/request_models"
	"github.com/turgut-nergin/tesodev/repository/models"
	"github.com/turgut-nergin/tesodev/repository/repo"
)

var UpdateCustomerHandler = func(r *repo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		if userId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user ID can not be empty"})
			return
		}

		var req request_models.Customer
		err := c.ShouldBindJSON(&req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		//#TODO: you must converted uuid to bson type :)
		customer := &models.Customer{
			Name:    req.Name,
			UserID:  userId,
			Email:   req.Email,
			Address: models.Address(req.Address),
		}

		customerR, err := r.Update(userId, customer)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusAccepted, customerR)
	}
}
