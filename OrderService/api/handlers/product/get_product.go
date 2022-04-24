package product

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
// 	"github.com/turgut-nergin/tesodev/database"
// )

// var GetOrdersHandler = func(database *database.Repository) func(context *gin.Context) {
// 	return func(c *gin.Context) {

// 		req, err := database.GetProducts()

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 			return
// 		}

// 		if req == nil {
// 			c.String(http.StatusNotFound, "Table is Empty")
// 			return
// 		}

// 		var products []response_models.Product
// 		for _, product := range req {
// 			products = append(products, response_models.Product{
// 				Name:      product.Name,
// 				ImageUrl:  product.ImageUrl,
// 				ProductId: product.ProductId,
// 			})
// 		}

// 		c.JSON(http.StatusOK, products)
// 	}
// }
