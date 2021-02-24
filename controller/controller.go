package controller

import (
  "fmt"
  "net/http"
  "context"
  "encoding/json"

  "trains/database"
  "trains/models"

  "go.mongodb.org/mongo-driver/bson"
)

func GetTrainsHandler(res http.ResponseWriter, req *http.Request) {
  var trains []models.Train

  collection := database.Client.Database("IRCTC").Collection("Trains")
  cursor, _ := collection.Find(context.TODO(), bson.D{})
  _ = cursor.All(context.TODO(), &trains)

  allTrains, _ := json.Marshal(trains)
  fmt.Println(string(allTrains))
  res.Write(allTrains)
}
