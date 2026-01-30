package connections

import (
	"context"
	"expensemanagement/internal/domain/resources/resourceserrormessages"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDB(ctx context.Context) (*mongo.Client, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return nil, fmt.Errorf(resourceserrormessages.UriIsEmpty)
	}

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf(resourceserrormessages.CouldNotConnectToDatabase)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf(resourceserrormessages.DatabaseConnectionIsNotStable)
	}

	return client, nil
}
