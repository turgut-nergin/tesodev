package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/customer"
	"github.com/turgut-nergin/tesodev/database"
)

func GetRouter(e *gin.Engine, r *database.Repository) {
	e.GET("/customer/get/customerId/:customerId", customer.GetCustomerByCustomerId(r))
	e.GET("/customer/bulk", customer.GetCustomers(r))
	e.POST("/customer/insert", customer.CreateCustomer(r))
	e.DELETE("/customer/delete/customerId/:customerId", customer.DeleteCustomerById(r))
	e.PUT("/customer/update/customerId/:customerId", customer.UpdateCustomerHandler(r))
	e.GET("/customer/validate/customerId/:customerId", customer.ValidateCustomer(r))
}
