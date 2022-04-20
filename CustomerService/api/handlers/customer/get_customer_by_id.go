package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
	"github.com/turgut-nergin/tesodev/repository/repo"
)

var GetCustomerByUserId = func(r *repo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		err := c.ShouldBind(userId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		req, err := r.GetByUserId(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is not found"})
			return
		}

		if req == nil {
			c.String(http.StatusNotFound, "Table is Empty")
			return
		}

		//#TODO: you must converted uuid to bson type :)
		customer := &response_models.Customer{
			Name:       req.Name,
			UserID:     req.UserID,
			Email:      req.Email,
			CreatedAdd: req.CreatedAdd,
			UpdatedAdd: req.UpdatedAdd,
			Address:    response_models.Address(req.Address),
		}

		c.JSON(http.StatusOK, customer)
	}
}
