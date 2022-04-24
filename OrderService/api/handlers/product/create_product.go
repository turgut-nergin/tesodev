package product

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/turgut-nergin/tesodev/api/handlers/request_models"
// 	"github.com/turgut-nergin/tesodev/database"
// 	"github.com/turgut-nergin/tesodev/database/models"
// )

// var CreateProductHandler = func(r *database.Repository) func(c *gin.Context) {
// 	return func(c *gin.Context) {

// 		var req *request_models.Product

// 		err := c.ShouldBindJSON(&req)

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		product := &models.Product{
// 			ImageUrl: req.ImageUrl,
// 			Name:     req.Name,
// 		}

// 		// validCustomer := customerValidation.Customer{
// 		// 	Customer: *customer,
// 		// }

// 		// err = validCustomer.Validate()

// 		// if err != nil {
// 		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		// 	return
// 		// }

// 		inserted, err := r.InsertProduct(product)

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, inserted)
// 	}
// }
