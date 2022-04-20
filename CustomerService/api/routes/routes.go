package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/customer"
	repositorty "github.com/turgut-nergin/tesodev/repository/repo"
)

func InitializeRoutes(e *gin.Engine, r *repositorty.Repository) {
	e.GET("/customers/userID/:userId", customer.GetCustomerByUserId(r))
	e.GET("/customers", customer.GetCustomers(r))
	e.POST("/customers", customer.CreateCustomer(r))
	e.DELETE("/customers/userId/:userId", customer.DeleteCustomerById(r))
}
