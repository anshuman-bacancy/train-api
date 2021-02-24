package database

import (
  "fmt"
  "time"
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

  chSize := 10
  ch := make(chan bool, chSize)

  now := time.Now()
  for _, train := range trains {
    ch <- true
    go func(train *models.Train) {
      _, _ = collection.InsertOne(context.TODO(), train)
    <-ch
    }(&train)
  }
  fmt.Println(time.Since(now))

  for x := 0; x < chSize; x++ {
    ch <- true
  }
}
