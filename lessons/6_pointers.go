package lessons

import "fmt"

func Pointers() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p) // *p here de-references to pointer to get the value
	fmt.Printf("\nThe value of i is: %v", i)

	// "Set the value of the memory location p is pointing at to 10"
	//*p = 10

	p = &i
	*p = 1
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)

	var k int32 = 2
	i = k
}

func withSlices() {
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4 // will change both slice + sliceCopy as slices are pointers to an underlying array
	fmt.Println(slice)
	fmt.Println(sliceCopy)
}

func withFunctions() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 is: %v", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

// We can use pointers to prevent having to copy large input params every time we call a function
func squarePointer(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 is: %v", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
