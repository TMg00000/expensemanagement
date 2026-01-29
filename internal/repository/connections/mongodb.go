package connections

import (
	"context"
	"expensemanagement/internal/configs"
	"expensemanagement/internal/domain/resources/resourceserrormessages"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type collection struct {
	configs.Config
}

func NewMongoDB(ctx context.Context) (*mongo.Client, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return nil, fmt.Errorf(resourceserrormessages.UriIsEmpty)
	}

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf(resourceserrormessages.CouldNotConnectToDatabase, err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf(resourceserrormessages.DatabaseConnectionIsNotStable, err)
	}

	return client, nil
}

func (col *collection) GetCollection(ctx context.Context, client *mongo.Client) error {
	settingsCol := client.Database("expensesdb").Collection("settings")

	err := settingsCol.FindOne(ctx, bson.M{
		"key": "EXPENSES_COLLECTION",
	}).Decode(&col.Config)
	if err != nil {
		return fmt.Errorf(resourceserrormessages.CouldNotFindThisKey, err)
	}

	return nil
}
