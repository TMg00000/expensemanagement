package main

import (
	"context"
	"expensemanagement/internal/configs"
	"expensemanagement/internal/domain/resources/resourceserrormessages"
	"expensemanagement/internal/http/handler"
	"expensemanagement/internal/repository"
	"expensemanagement/internal/repository/connections"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	startConfig := configs.StartConfig()
	returnFatalError(resourceserrormessages.EnvironmentConfigInitError, startConfig)

	db := startDataBase()
	defer func() {
		if discErr := db.Database().Client().Disconnect(context.Background()); discErr != nil {
			log.Fatal(discErr)
		}
	}()

	expensesRepository := repository.NewExpensesRepository(db)

	controller := &handler.ExpensesServices{
		Services: expensesRepository,
	}

	r := mux.NewRouter()
	r.HandleFunc("/expenses", controller.RegisterExpenses).Methods("POST")

	err := http.ListenAndServe(":9437", r)
	returnFatalError(err.Error(), err)
}

func startDataBase() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := connections.NewMongoDB(ctx)
	returnFatalError(resourceserrormessages.DatabaseConnectionError, err)

	return client.Database("expensesdb").Collection(configs.Env.EXPENSES_COL)
}

func returnFatalError(errormessage string, err error) {
	if err != nil {
		log.Fatal(errormessage, err)
	}
}
