package main

import (
	"flag"
	"log"
	"net/http"

	"trains/controller"
	"trains/database"
)

func main() {
	initErr := database.InitServer()
	if initErr != nil {
		log.Fatal(initErr)
		panic(initErr)
	}

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
	http.HandleFunc("/getTrainsFilter", controller.GetTrainsFilterHandler)

	log.Println("Server started at", port)
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	http.Handle("/static/", fs)

	http.ListenAndServe(port, nil)
}
