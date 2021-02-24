package database

import (
  "os"
  "encoding/csv"

	"trains/models"
  "trains/utils"
)

var FileName = "All_Indian_Trains.csv" 

// Reads csv and parses into []models.Train
func IRCTCdata() []models.Train {
  trainData, err := os.Open(FileName)
  utils.CheckError(err)
  csvReader := csv.NewReader(trainData)
  trainRecords, _ := csvReader.ReadAll()

	var allTrains []models.Train

	for _, trainRecord := range trainRecords[1:] {
		var train models.Train

		train.TrainNo = trainRecord[1]
		train.TrainName = trainRecord[2]
		train.TrainStarts = trainRecord[3]
		train.TrainEnds = trainRecord[4]

		allTrains = append(allTrains, train)
	}

	return allTrains
}
