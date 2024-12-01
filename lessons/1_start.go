package lessons

import (
	"errors"
	"fmt"
)

func Start() {
	var intNum int64
	fmt.Println(intNum)

	var1, var2 := 1, 2
	fmt.Println(var1 + var2)

	const immutable int = 12
	fmt.Println(immutable)

	printMe("hello world")

	var result, remainder, err = intDivision(11, 2)

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result is: %v", result)
	default:
		fmt.Printf("The result is: %v with remainder: %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("Remainder is: 0")
	case 1:
		fmt.Println("Remainder is: 1")
	default:
		fmt.Println("Remainder > 1")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}

	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, err
}
