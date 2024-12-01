package lessons

import (
	"encoding/json"
	"fmt"
	"os"
)

func Generics() {
	sumSlice([]float32{1, 2, 3})
	sumSlice([]int{4, 5, 6})
	isEmpty([]string{})
	isEmpty([]string{"string"})
	isEmpty([]int{})

	loadData()
	makeCars()
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumIntSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumFloat32Slice(slice []float32) float32 {
	var sum float32
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumFloat64Slice(slice []float64) float64 {
	var sum float64
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func loadData() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%v", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	var data, _ = os.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

type otherGasEngine struct {
	gallons float32
	mpg     float32
}

type otherElectricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T otherGasEngine | otherElectricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func makeCars() {
	var gasCar = car[otherGasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: otherGasEngine{
			gallons: 12.4,
			mpg:     40.0,
		},
	}

	fmt.Println(gasCar)

	var electricCar = car[otherElectricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: otherElectricEngine{
			kwh:   57.5,
			mpkwh: 4.17,
		},
	}
	fmt.Println(electricCar)

}
