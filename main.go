package main

import (
  "os"
  "log"
  "net/http"

  "trains/routes"
  "trains/models"
  "trains/data"
  "trains/database"
)

func main() {
  trainRecords := data.IRCTCdata()
  var allTrains []models.Train

  // parse train data into struct
  for _, data := range trainRecords[1:] {
    var train models.Train

    train.TrainNo = data[1]
    train.TrainName = data[2]
    train.TrainStarts = data[3]
    train.TrainEnds = data[4]

    allTrains = append(allTrains, train)
  }

  database.InitServer()
  //database.BulkInsert(allTrains)

  //routes
  http.HandleFunc("/getTrains", routes.GetTrainsHandler)

  port := string(":" + os.Args[1])

  log.Println("Server started at", port)
  // file server for images
  fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
  http.Handle("/static/", fs)

  http.ListenAndServe(port, nil)
}
