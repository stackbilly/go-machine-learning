package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/knn"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//load dataset
	data, err := base.ParseCSVToInstances("/datasets/insurance.csv", true)

	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("data:\n%s", data)

	trainData, testData := base.InstancesTrainTestSplit(data, 0.80)
	knnRegr := knn.NewKnnRegressor("euclidean")

	err = knnRegr.Fit(trainData)
}
