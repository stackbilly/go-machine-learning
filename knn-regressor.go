package main

import (
	"fmt"
	mongoimport "github.com/Livingstone-Billy/mongo-import"
	"strings"
)

type Insurance struct {
	Age      int
	Sex      string
	Bmi      float64
	Children int
	Smoker   int
	Region   string
	Charges  float64
}

func loadDataset(filename string) ([]Insurance, error) {
	records, err := mongoimport.CSVReader("datasets/insurance.csv")
	if err != nil {
		return nil, err
	}
	//parse csv data into a slice of house objects
	insurances := []Insurance{}
	for _, line :=range strings.Split()
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	//
	////load dataset
	//data, err := base.ParseCSVToInstances("/datasets/insurance.csv", true)
	//
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
	//fmt.Printf("data:\n%s", data)
	//
	//trainData, testData := base.InstancesTrainTestSplit(data, 0.80)
	//knnRegr := knn.NewKnnRegressor("euclidean")
	//
	//err = knnRegr.Fit(trainData)

	records, err := mongoimport.CSVReader("datasets/insurance.csv")
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Records:\n%s", records)
}
