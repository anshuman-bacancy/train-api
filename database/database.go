package database

import (
  "context"

  "trains/models"
  "trains/utils"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var clientOps *options.ClientOptions
var Client *mongo.Client
var err error
var DbConnURI = "mongodb://localhost:27017"
var DbName = "IRCTC"
var CollectionName = "Trains"

func InitServer() {
  clientOps = options.Client().ApplyURI(DbConnURI)
  Client, err = mongo.Connect(context.TODO(), clientOps)
  utils.CheckError(err)
}

func BulkInsert(trains []models.Train) {
  collection := Client.Database(DbName).Collection(CollectionName)
  for _, train := range trains {
    _, _ = collection.InsertOne(context.TODO(), train)
  }
}
