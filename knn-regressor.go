package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/knn"
	"io/ioutil"
	"strconv"
	"strings"
)

type Insurance struct {
	Age      float64
	Sex      float64
	Bmi      float64
	Children float64
	Smoker   float64
	Region   float64
	Charges  float64
}

func loadDataset(filename string) ([]Insurance, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
		return nil, err
	}
	var insurances []Insurance
	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Split(line, ",")
		insurance := Insurance{}
		insurance.Age, _ = strconv.ParseFloat(parts[0], 64)
		insurance.Sex, _ = strconv.ParseFloat(parts[1], 64)
		insurance.Bmi, _ = strconv.ParseFloat(parts[2], 64)
		insurance.Children, _ = strconv.ParseFloat(parts[3], 64)
		insurance.Smoker, _ = strconv.ParseFloat(parts[4], 64)
		insurance.Region, _ = strconv.ParseFloat(parts[5], 64)
		insurance.Charges, _ = strconv.ParseFloat(parts[6], 64)
		insurances = append(insurances, insurance)
	}
	return insurances, err
}

func main() {
	//data, err := base.ParseCSVToInstances("datasets/insurance.csv", true)
	data, err := loadDataset("datasets/insurance.csv")
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Data:\n%v", data)

	var age []float64
	for _, val := range data {
		age = append(age, val.Age)
	}

	var bmi []float64
	for _, val := range data {
		bmi = append(bmi, val.Bmi)
	}

	var sex []float64
	for _, val := range data {
		sex = append(sex, val.Sex)
	}

	var targets []float64
	for _, val := range data {
		targets = append(targets, val.Charges)
	}

	var trainData []float64
	trainData = append(trainData, age...)
	trainData = append(trainData, sex...)
	trainData = append(trainData, bmi...)

	knnReg := knn.NewKnnRegressor("euclidean")

	knnReg.Fit(trainData, targets, len(trainData), 3)

	gender, _ := strconv.ParseFloat("female", 64)

	var newInstance = []float64{34, gender, 34.5}

	predictions := knnReg.Predict(newInstance, 3)
}
