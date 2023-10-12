package main

import (
	"fmt"
	mongoimport "github.com/Livingstone-Billy/mongo-import"
	"github.com/sjwhitworth/golearn/knn"
	"gonum.org/v1/gonum/mat"
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
	//data, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	panic(err)
	//	return nil, err
	//}
	var insurances []Insurance

	records, err := mongoimport.CSVReader("datasets/insurance.csv")
	if err != nil {
		panic(err)
		return nil, err
	}
	//fmt.Println(records[0])

	for i := 1; i < len(records); i++ {
		insurance := Insurance{}
		insurance.Age, _ = strconv.ParseFloat(records[i][0], 64)
		insurance.Sex, _ = strconv.ParseFloat(records[i][1], 64)
		insurance.Bmi, _ = strconv.ParseFloat(records[i][2], 64)
		insurance.Children, _ = strconv.ParseFloat(records[i][3], 64)
		insurance.Smoker, _ = strconv.ParseFloat(records[i][4], 64)
		insurance.Region, _ = strconv.ParseFloat(records[i][5], 64)
		insurance.Charges, _ = strconv.ParseFloat(records[i][6], 64)
		insurances = append(insurances, insurance)
	}

	//for _, val := range data {
	//	parts := strings.Split(string(val), ",")
	//	insurance := Insurance{}
	//	insurance.Age, _ = strconv.ParseFloat(parts[0], 64)
	//	insurance.Sex, _ = strconv.ParseFloat(parts[1], 64)
	//	insurance.Bmi, _ = strconv.ParseFloat(parts[2], 64)
	//	insurance.Children, _ = strconv.ParseFloat(parts[3], 64)
	//	insurance.Smoker, _ = strconv.ParseFloat(parts[4], 64)
	//	insurance.Region, _ = strconv.ParseFloat(parts[5], 64)
	//	insurance.Charges, _ = strconv.ParseFloat(parts[6], 64)
	//	insurances = append(insurances, insurance)
	//}
	//fmt.Printf("insurances:\n%s", insurances)
	return insurances, err
}

func main() {
	data, err := mongoimport.CSVReader("datasets/insurance.csv")
	if err != nil {
		panic(err)
		return
	}

	size := len(data)
	var age []float64
	for i := 1; i < size; i++ {
		val, _ := strconv.ParseFloat(data[i][0], 64)
		age = append(age, val)
	}
	var sex []float64
	for j := 1; j < size; j++ {
		if strings.Compare("male", data[j][1]) == 0 {
			sex = append(sex, 1)
		} else {
			sex = append(sex, 0)
		}
	}

	var bmi []float64
	for j := 1; j < size; j++ {
		val, _ := strconv.ParseFloat(data[j][2], 64)
		bmi = append(bmi, val)
	}

	var targets []float64
	for k := 1; k < size; k++ {
		val, _ := strconv.ParseFloat(data[k][6], 64)
		targets = append(targets, val)
	}
	//
	//var trainData []float64
	//trainData = append(trainData, age...)
	//trainData = append(trainData, sex...)
	//trainData = append(trainData, bmi...)

	knnReg := knn.NewKnnRegressor("euclidean")

	knnReg.Fit(age, targets, len(age), 1)

	var newInstance = []float64{34, 1, 34.5}

	dense := mat.NewDense(1, 3, newInstance)

	predictions := knnReg.Predict(dense, 3)
	fmt.Println(predictions)
}
