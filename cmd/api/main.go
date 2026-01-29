package main

import (
	"context"
	"expensemanagement/internal/configs"
	"expensemanagement/internal/domain/resources/resourceserrormessages"
	"expensemanagement/internal/repository/connections"
	"log"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	startConfig := configs.StartConfig()
	returnFatalError(resourceserrormessages.EnvironmentConfigInitError, startConfig)

	startDataBase()

	r := mux.NewRouter()
	r.HandleFunc("/", nil)
}

func startDataBase() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := connections.NewMongoDB(ctx)
	returnFatalError(resourceserrormessages.DatabaseConnectionError, err)

	return client.Database("expensedb").Collection(configs.Env.EXPENSES_COL)
}

func returnFatalError(errormessage string, err error) {
	if err != nil {
		log.Fatal(errormessage, err)
	}
}
