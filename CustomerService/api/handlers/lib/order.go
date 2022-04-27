package lib

import (
	"time"

	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
	"github.com/turgut-nergin/tesodev/database/models"
)

func ResponseAssign(customer *models.Customer) *response_models.Customer {
	return &response_models.Customer{
		Name:       customer.Name,
		CustomerId: customer.CustomerId,
		Email:      customer.Email,
		CreatedAt:  time.Unix(customer.CreatedAt, 0),
		UpdatedAt:  time.Unix(customer.UpdatedAt, 0),
		Address:    response_models.Address(customer.Address),
	}

}
