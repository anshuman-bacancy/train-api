package main

import (
  "log"
  "flag"
  "net/http"

  "trains/controller"
  "trains/database"
)

func main() {
  database.InitServer()

  processcsv := flag.Bool("processcsv", false, "")
  flag.Parse()

  if *processcsv {
    allTrains := database.IRCTCdata()
    database.BulkInsert(allTrains)
  }

  var port string
  if len(flag.Args()) == 1 {
    port = string(":" + flag.Args()[0])
  } else {
    port = ":8000"
  }

  //handling routes
  http.HandleFunc("/getTrains", controller.GetTrainsHandler)

  log.Println("Server started at", port)
  fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
  http.Handle("/static/", fs)

  http.ListenAndServe(port, nil)
}
