package lessons

import "fmt"

type exampleEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	other
}

type owner struct {
	name string
}

type other struct {
	other string
}

func StructsAndInterfaces() {
	var myExampleEngine exampleEngine
	fmt.Printf("On instantiation with no values, we get a default value. Object: { %v, %v }", myExampleEngine.mpg, myExampleEngine.gallons)

	myExampleEngine.mpg = 20
	fmt.Printf("After setting: { %v }", myExampleEngine.mpg)

	var myExampleEngine2 = exampleEngine{25, 15, owner{"Jeff"}, other{"other"}}
	fmt.Println("Create a new exampleEngine")
	fmt.Println(myExampleEngine2.mpg, myExampleEngine2.gallons, myExampleEngine2.ownerInfo, myExampleEngine2.other)

	var dieselEngine = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println("Inline anonymous struct")
	fmt.Printf("dieselEngine { %v, %v }", dieselEngine.mpg, dieselEngine.gallons)

	fmt.Println("Miles Left")
	var myGasEngine gasEngine = gasEngine{25, 15}
	fmt.Printf("Total miles left in tank: %v", myGasEngine.milesLeft())

	var myElectricEngine = electricEngine{35, 10}
	canMakeIt(myGasEngine, 100)
	canMakeIt(myElectricEngine, 100)
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Will make it")
	} else {
		fmt.Println("Will NOT make it")
	}
}
