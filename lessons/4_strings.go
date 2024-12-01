package lessons

import (
	"fmt"
	"strings"
)

func Strings() {
	var myString = "Résumé"
	var indexed = myString[0]

	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}
	// Length is the number of BYTES in the underlying representation of our strings
	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	var myRuneString []rune = []rune(myString)
	var indexE = myString[1]      // X
	var indexE2 = myRuneString[1] // returns 233 as expected
	fmt.Printf("%v != %v", indexE, indexE2)

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"a", "b", "c", "d", "e", "f"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)

}
