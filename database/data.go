package database

import (
	"encoding/csv"
	"log"
	"os"

	"trains/models"
)

var fileName = "rail-data.csv"

// IRCTCdata reads csv and parses into []models.Train
func IRCTCdata() []models.Train {
	trainData, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(trainData)
	trainRecords, _ := csvReader.ReadAll()

	var allTrains []models.Train

	for _, trainRecord := range trainRecords {
		var train models.Train

		train.TrainNo = trainRecord[0]
		train.TrainName = trainRecord[1]
		train.SEQ = trainRecord[2]
		train.Code = trainRecord[3]
		train.StName = trainRecord[4]
		train.ATime = trainRecord[5]
		train.DTime = trainRecord[6]
		train.Distance = trainRecord[7]
		train.SS = trainRecord[8]
		train.SSname = trainRecord[9]
		train.Ds = trainRecord[10]
		train.DsName = trainRecord[11]

		allTrains = append(allTrains, train)
	}

	return allTrains
}
