package create

import (
	"time"
)

type expenses struct {
	Name        string
	Description string
	Value       float32
	DueDate     time.Time
}
