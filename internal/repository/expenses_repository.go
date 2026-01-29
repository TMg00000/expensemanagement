package repository

import (
	"context"
	"expensemanagement/internal/domain/request"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExpensesRepository struct {
	collection *mongo.Collection
}

func NewExpensesRepository(c *mongo.Collection) *ExpensesRepository {
	return &ExpensesRepository{collection: c}
}

func (c *ExpensesRepository) Create(e request.Expenses) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := bson.M{
		"name":        e.Name,
		"description": e.Description,
		"value":       e.Value,
		"duedate":     e.DueDate,
	}

	_, err := c.collection.InsertOne(ctx, doc)
	return err
}
