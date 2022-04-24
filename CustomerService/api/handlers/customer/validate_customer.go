package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/turgut-nergin/tesodev/database"
)

var ValidateCustomer = func(r *database.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {

		customerId := c.Param("customerId")
		println(customerId)
		_, err := uuid.Parse(customerId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		isExist, err := r.IdIsExist(customerId)
		println(isExist)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusOK, isExist)
	}
}
