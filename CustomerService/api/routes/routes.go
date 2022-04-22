package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/customer"
	repositorty "github.com/turgut-nergin/tesodev/repository/repo"
)

func InitializeRoutes(e *gin.Engine, r *repositorty.Repository) {
	e.GET("/get/customers/userID/:userId", customer.GetCustomerByUserId(r))
	e.GET("/bulk/customers", customer.GetCustomers(r))
	e.POST("/insert/customer", customer.CreateCustomer(r))
	e.DELETE("/delete/customers/userId/:userId", customer.DeleteCustomerById(r))
	e.PUT("/update/customers/userId/:userId", customer.UpdateCustomerHandler(r))
}
