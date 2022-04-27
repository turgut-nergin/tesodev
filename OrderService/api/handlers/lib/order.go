package lib

import (
	"time"

	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
	"github.com/turgut-nergin/tesodev/database/models"
)

func ResponseAssign(order *models.Order) *response_models.Order {

	return &response_models.Order{
		OrderId:    order.OrderId,
		CustomerId: order.CustomerId,
		Quantity:   order.Quantity,
		Price:      order.Price,
		Status:     order.Status,
		Address:    response_models.Address(order.Address),
		Product:    response_models.Product(order.Product),
		CreatedAt:  time.Unix(order.CreatedAt, 0),
		UpdatedAt:  time.Unix(order.UpdatedAt, 0),
	}

}
