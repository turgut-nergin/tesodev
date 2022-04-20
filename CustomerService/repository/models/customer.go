package models

import (
	"time"
)

type Customer struct {
	UserID     string    `bson:"userId"`
	Name       string    `bson:"name"`
	Email      string    `bson:"email"`
	Address    Address   `bson:"adress"`
	CreatedAdd time.Time `bson:"createdAdd"`
	UpdatedAdd time.Time `bson:"updatedAdd"`
}
