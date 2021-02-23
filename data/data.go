package data

import (
  "os"
  "encoding/csv"
  "trains/utils"
)

func IRCTCdata() [][]string {
  trainData, err := os.Open("All_Indian_Trains.csv")
  utils.CheckError(err)
  csvReader := csv.NewReader(trainData)
  data, _ := csvReader.ReadAll()

  return data
}
