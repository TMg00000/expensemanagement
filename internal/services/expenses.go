package services

import (
	"expensemanagement/internal/domain/request"
)

type Expenses interface {
	Create(request.Expenses) error
}
