package database

import (
  "log"
  "fmt"
  "time"
  "context"

  "trains/models"
  "trains/utils"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var (
  clientOps *options.ClientOptions
  Client *mongo.Client
  clientErr error
  config utils.Config
  configErr error
  insertErr error
)


func init() {
  config, configErr = utils.LoadConfig(".")
  if configErr != nil {
    log.Fatal(configErr)
    panic(configErr)
  }
}

func InitServer() error {
  clientOps = options.Client().ApplyURI(config.ConnString)
  Client, clientErr = mongo.Connect(context.TODO(), clientOps)
  return clientErr
}

func BulkInsert(trains []models.Train) error {
  collection := Client.Database(config.DbName).Collection(config.Collection)

  chSize := 10
  ch := make(chan bool, chSize)

  now := time.Now()
  for _, train := range trains {
    ch <- true
    go func(train *models.Train) {
      _, insertErr = collection.InsertOne(context.TODO(), train)
    <-ch
    }(&train)
  }
  fmt.Println("Time elapsed.. ", time.Since(now))

  for x := 0; x < chSize; x++ {
    ch <- true
  }
  return insertErr
}
