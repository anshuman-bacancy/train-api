package main

import (
  "log"
  "flag"
  "net/http"

  "trains/controller"
  "trains/models"
  "trains/database"
)

func main() {
  trainRecords := database.IRCTCdata()
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

  bulkinsert := flag.Bool("bulkinsert", false, "")
  flag.Parse()

  var port string
  if len(flag.Args()) == 1 {
    port = string(":" + flag.Args()[0])
  } else {
    port = ":8000"
  }

  database.InitServer()
  if *bulkinsert {
	  database.BulkInsert(allTrains)
  }

  //handling routes
  http.HandleFunc("/getTrains", controller.GetTrainsHandler)

  log.Println("Server started at", port)
  fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
  http.Handle("/static/", fs)

  http.ListenAndServe(port, nil)
}
