package request_models

import (
	"time"
)

type Customer struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Address    Address   `json:"address"`
	CreatedAdd time.Time `json:"createdAdd"`
	UpdatedAdd time.Time `json:"updatedAdd"`
}
