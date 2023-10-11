package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func KnnClassifier(dataset string) (base.FixedDataGrid, map[string]map[string]int) {
	data, err := base.ParseCSVToInstances(dataset, true)

	if err != nil {
		panic(err)
		return nil, nil
	}

	//summary of the
	fmt.Printf("Data:\n%s", data)

	knnModel := knn.NewKnnClassifier("euclidean", "linear", 2)

	trainData, testData := base.InstancesTrainTestSplit(data, 0.50)

	err = knnModel.Fit(trainData)
	if err != nil {
		panic(err)
		return nil, nil
	}
	predictions, err := knnModel.Predict(testData)
	if err != nil {
		panic(err)
		return nil, nil
	}
	confusionMatrix, _ := evaluation.GetConfusionMatrix(predictions, testData)

	return predictions, confusionMatrix
}

// simple iris knn classifier
func main() {
	predictions, confusionMat := KnnClassifier("Iris.csv")
	fmt.Printf("Predictions:\n%s", predictions)
	fmt.Printf("\nConfusion Matrix:\n%s", evaluation.GetSummary(confusionMat))
}

/*sample output
Confusion Matrix:
Reference Class True Positives  False Positives True Negatives  Precision      Recall                                                                        F
1 Score
--------------- --------------  --------------- --------------  ---------      ------                                                                        -
-------
Iris-versicolor 21              1               52              0.9545         0.9130                                                                        0
.9333
Iris-virginica  26              2               48              0.9286         1.0000                                                                        0
.9630
Iris-setosa     26              0               49              1.0000         0.9630                                                                        0
.9811
Overall accuracy: 0.9605*/
