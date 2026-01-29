package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type collection struct {
	collection *mongo.Collection
}

func Create() {
	doc := bson.M{}
}
