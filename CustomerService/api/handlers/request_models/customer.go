package request_models

import (
	"time"
)

type Customer struct {
	UserID     string    `json:"userId"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Address    Address   `json:"adress"`
	CreatedAdd time.Time `json:"createdAdd"`
	UpdatedAdd time.Time `json:"updatedAdd"`
}
