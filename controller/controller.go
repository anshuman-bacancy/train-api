package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"trains/database"
	"trains/models"

	"go.mongodb.org/mongo-driver/bson"
)

// GetTrainsHandler shows all trains
func GetTrainsHandler(res http.ResponseWriter, req *http.Request) {
	var trains []models.Train

	collection := database.Client.Database("trains_db").Collection("trains")
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	_ = cursor.All(context.TODO(), &trains)
	fmt.Println("hi")

	json.NewEncoder(res).Encode(trains[1:101])
}

//bubble sort

func sort(trains []models.Train) []models.Train { //bubble sort
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

func subTime(time1, time2 string) int {
	time1 = strings.Replace(time1, ":", "", -1)
	time2 = strings.Replace(time2, ":", "", -1)
	time1Int, _ := strconv.Atoi(time1)
	time2Int, _ := strconv.Atoi(time2)
	if time1Int > time2Int {
		return time1Int - time2Int
	}
	return time2Int - time1Int
}

// GetTrainsFilterHandler filters trains
func GetTrainsFilterHandler(res http.ResponseWriter, req *http.Request) {
	source := req.URL.Query().Get("src")
	dst := req.URL.Query().Get("dst")

	collection := database.Client.Database("trains_db").Collection("trains")
	cur, _ := collection.Find(context.TODO(), bson.M{"code": source})
	var station1 []models.Train
	if err := cur.All(context.TODO(), &station1); err != nil {
		log.Fatal(err)
	}

	cur, _ = collection.Find(context.TODO(), bson.M{"code": dst})
	var station2 []models.Train
	if err := cur.All(context.TODO(), &station2); err != nil {
		log.Fatal(err)
	}

	var sourceTrains []models.Train

	count := 0
	for _, src := range station1 {
		for _, dst := range station2 {
			if src.TrainNo == dst.TrainNo {
				srcSeq, _ := strconv.Atoi(src.SEQ)
				dstSeq, _ := strconv.Atoi(dst.SEQ)
				if srcSeq < dstSeq {
					src.TimeDiff = subTime(src.ATime, dst.ATime)
					sourceTrains = append(sourceTrains, src)
					count++
				}
			}
		}
	}
	fmt.Println(len(sourceTrains))

	sortedTrains := sort(sourceTrains)

	// sort.Slice(sourceTrains, func(i, j int) bool {
	// 	return sourceTrains[i].TimeDiff < sourceTrains[j].TimeDiff
	// })

	fmt.Println(count)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(sortedTrains)
}
