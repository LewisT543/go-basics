package lessons

import "fmt"

func Collections() {
	printArraysAndLocations()
	createArrays()
	maps()
	forLoops()
	shorthand()
}

func printArraysAndLocations() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// &variable for memory location of variable
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
}

func createArrays() {
	var intArr = [3]int32{1, 2, 3}
	intArr2 := [...]int32{1, 2, 3}
	fmt.Println(intArr)
	fmt.Println(intArr2)

	// dynamic array sizing
	var intSlice = []int32{4, 5, 6}
	fmt.Printf("The length is: %v the capacity is %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is: %v the capacity is %v", len(intSlice), cap(intSlice))

	var intSlice2 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("The length is: %v the capacity is %v", len(intSlice), cap(intSlice))

	var intSlice3 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
}

func maps() {
	var myMap = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["George"])

	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

}

func forLoops() {
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	var intArr = []uint32{1, 2, 3}
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	var i = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i += 1
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func shorthand() {
	var n = 10
	n++
	n--
	n += 1
	n -= 1
	n *= 2
	n /= 2
}
