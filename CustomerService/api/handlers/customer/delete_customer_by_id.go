package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/repository/repo"
)

var DeleteCustomerById = func(r *repo.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		if len(userId) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id can not be empty!"})
			return
		}

		err := r.Delete(userId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is not found"})
			return
		}

		c.JSON(http.StatusOK, true)
	}
}
