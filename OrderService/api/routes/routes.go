package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/order"
	"github.com/turgut-nergin/tesodev/database"
)

func InitializeRoutes(e *gin.Engine, r *database.Repository) {
	e.POST("/insert/order/customerId/:customerId", order.CreateOrderHandler(r))
	e.PUT("/update/orders/orderId/:orderId", order.UpdateOrderHandler(r))
	e.DELETE("/delete/orders/orderId/:orderId", order.DeleteOrderById(r))
	e.GET("/bulk/orders", order.GetOrdersHandler(r))
	e.GET("/bulk/customer/orders/customerId/:customerId", order.GetOrdersByCustomerId(r))
	e.GET("/get/orders/orderId/:orderId", order.GetOrderById(r))
	e.PUT("/update/order/status/orderId/:orderId", order.ChangeOrderStatus(r))
}
