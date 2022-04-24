package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/turgut-nergin/tesodev/api/handlers/order"
	"github.com/turgut-nergin/tesodev/database"
)

func GetRouter(e *gin.Engine, r *database.Repository) {
	e.POST("/order/insert/customerId/:customerId", order.CreateOrderHandler(r))
	e.PUT("/order/update/orderId/:orderId", order.UpdateOrderHandler(r))
	e.DELETE("/order/delete/orderId/:orderId", order.DeleteOrderById(r))
	e.GET("/order/bulk", order.GetOrdersHandler(r))
	e.GET("/order/bulk/customer/customerId/:customerId", order.GetOrdersByCustomerId(r))
	e.GET("/order/get/orderId/:orderId", order.GetOrderById(r))
	e.PUT("/order/update/status/orderId/:orderId", order.ChangeOrderStatus(r))
}
