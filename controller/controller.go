package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"trains/database"
	"trains/models"
	"trains/utils"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	config    models.Config
	configErr error
)

func init() {
	config, configErr = utils.LoadConfig(".")
	if configErr != nil {
		log.Fatal(configErr)
		panic(configErr)
	}
}

// GetTrainsHandler shows all trains
func GetTrainsHandler(res http.ResponseWriter, req *http.Request) {
	var trains []models.Train

	collection := database.Client.Database(config.DbName).Collection(config.Collection)
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	_ = cursor.All(context.TODO(), &trains)
	json.NewEncoder(res).Encode(trains[1:101])
}

// GetTrainsFilterHandler filters trains according to query parameters
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
					src.TimeDiff = utils.TrainTimeDiff(src.ATime, dst.ATime)
					sourceTrains = append(sourceTrains, src)
					count++
				}
			}
		}
	}
	sortedTrains := utils.SortTrains(sourceTrains)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(sortedTrains)
}
