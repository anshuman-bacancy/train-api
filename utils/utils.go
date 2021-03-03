package utils

import (
	"strconv"
	"strings"
	"trains/models"
)

// SortTrains sorts according to time
func SortTrains(trains []models.Train) []models.Train {
	n := len(trains)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if trains[i-1].TimeDiff > trains[i].TimeDiff {
				trains[i], trains[i-1] = trains[i-1], trains[i]
				swapped = true
			}
		}
	}
	return trains
}

// TrainTimeDiff returns time difference of train arrival time
func TrainTimeDiff(time1, time2 string) int {
	time1 = strings.Replace(time1, ":", "", -1)
	time2 = strings.Replace(time2, ":", "", -1)
	time1Int, _ := strconv.Atoi(time1)
	time2Int, _ := strconv.Atoi(time2)
	if time1Int > time2Int {
		return time1Int - time2Int
	}
	return time2Int - time1Int
}
