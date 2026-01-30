package repository

import (
	"context"
	"expensemanagement/internal/domain/request"
	"expensemanagement/internal/domain/resources/resourceserrormessagesrepository"
	"fmt"
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
	if err != nil {
		return err
	}

	return nil
}

func (c *ExpensesRepository) GetAll() ([]request.Expenses, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := c.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var allExpenses []request.Expenses
	if err = cursor.All(ctx, &allExpenses); err != nil {
		return nil, err
	}

	return allExpenses, nil
}

func (c *ExpensesRepository) UpdateById(e request.Expenses) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": e.Id}
	doc := bson.M{
		"$set": bson.M{
			"name":        e.Name,
			"description": e.Description,
			"value":       e.Value,
			"duedate":     e.DueDate,
		},
	}

	result, err := c.collection.UpdateOne(ctx, filter, doc)
	if result.MatchedCount < 1 {
		return fmt.Errorf(resourceserrormessagesrepository.NoDocumentWasFound)
	}
	if err != nil {
		return err
	}
	return nil
}

func (c *ExpensesRepository) DeleteById(e request.Expenses) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": e.Id}

	_, err := c.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (c *ExpensesRepository) DeleteAll(e request.Expenses) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}
	return nil
}
