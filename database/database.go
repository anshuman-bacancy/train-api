package database

import (
  "context"
  "trains/models"
  "trains/utils"
  //"go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var clientOps *options.ClientOptions
var Client *mongo.Client
var err error

func InitServer() {
  clientOps = options.Client().ApplyURI("mongodb://localhost:27017")
  Client, err = mongo.Connect(context.TODO(), clientOps)
  utils.CheckError(err)
}

func BulkInsert(trains []models.Train) {
  collection := Client.Database("IRCTC").Collection("Trains")
  for _, train := range trains {
    //_, _ = collection.UpdateOrReplace(context.TODO(), train)
    _, _ = collection.InsertOne(context.TODO(), train)
  }
}



