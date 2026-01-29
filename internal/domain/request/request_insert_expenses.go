package request

import (
	"time"
)

type Expenses struct {
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Value       float32   `json:"value" bson:"value"`
	DueDate     time.Time `json:"duedate" bson:"duedate"`
}
