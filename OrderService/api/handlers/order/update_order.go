package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/request_models"
	"github.com/turgut-nergin/tesodev/database"
	"github.com/turgut-nergin/tesodev/database/models"
)

var UpdateOrderHandler = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		orderId := c.Param("orderId")

		if orderId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user ID can not be empty"})
			return
		}

		var req *request_models.Order
		err := c.ShouldBindJSON(&req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// validRequest := requestValidation.Customer{
		// 	Customer: req,
		// }

		// err = validRequest.Validate()

		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		order := &models.Order{
			OrderId:  orderId,
			Quantity: req.Quantity,
			Price:    req.Price,
			Status:   req.Status,
			Address:  models.Address(req.Address),
			Product:  models.Product(req.Product),
		}

		_, err = r.Update(orderId, order)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusAccepted, true)
	}
}
