package request

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expenses struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Value       float32            `json:"value" bson:"value"`
	DueDate     time.Time          `json:"duedate" bson:"duedate"`
}
