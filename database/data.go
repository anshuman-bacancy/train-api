package database

import (
  "os"
  "encoding/csv"

  "trains/utils"
)

var FileName = "All_Indian_Trains.csv" 

func IRCTCdata() [][]string {
  trainData, err := os.Open(FileName)
  utils.CheckError(err)
  csvReader := csv.NewReader(trainData)
  data, _ := csvReader.ReadAll()

  return data
}
