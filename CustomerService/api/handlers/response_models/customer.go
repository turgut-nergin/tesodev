package response_models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Customer struct {
	ID         bson.ObjectId `bson:"_id" json:"id,omitempty"`
	UserID     string        `bson:"userId"`
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Address    Address       `json:"adress"`
	CreatedAdd time.Time     `json:"createdAdd"`
	UpdatedAdd time.Time     `json:"updatedAdd"`
}
