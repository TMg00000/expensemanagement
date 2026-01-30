package services

import (
	"expensemanagement/internal/domain/request"
)

type Expenses interface {
	Create(request.Expenses) error
	GetAll() ([]request.Expenses, error)
	UpdateById(request.Expenses) error
	DeleteById(request.Expenses) error
	DeleteAll(request.Expenses) error
}
