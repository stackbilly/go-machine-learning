package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	data, err := base.ParseCSVToInstances("Iris.csv", true)
	if err != nil {
		panic(err)
		return
	}
	//a summary of the data
	fmt.Println(data)
	knnModel := knn.NewKnnClassifier("euclidean", "linear", 2)

	trainD, testD := base.InstancesTrainTestSplit(data, 0.70)
	err = knnModel.Fit(trainD)
	if err != nil {
		panic(err)
		return
	}
	predictions, err := knnModel.Predict(testD)
	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("Predictions:\n%s", predictions)
	confusionMat, err := evaluation.GetConfusionMatrix(predictions, testD)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Summary\n%s", evaluation.GetSummary(confusionMat))
}
