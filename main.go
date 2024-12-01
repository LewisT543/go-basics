package main

import (
	"fmt"
	"github.com/LewisT543/go-basics/lessons"
)

func main() {
	fmt.Println("<<<<Start>>>>")
	lessons.Start()

	fmt.Println("<<<<Collections>>>>")
	lessons.Collections()

	fmt.Println("<<<<SpeedSlice>>>>")
	lessons.SpeedSlice()

	fmt.Println("<<<<Strings>>>>")
	lessons.Strings()

	fmt.Println("<<<<StructsAndInterfaces>>>>")
	lessons.StructsAndInterfaces()

	fmt.Println("<<<<Pointers>>>>")
	lessons.Pointers()

	fmt.Println("<<<<GoRoutines>>>>")
	lessons.GoRoutines()

	fmt.Println("<<<<Channels>>>>")
	lessons.Channels()

	fmt.Println("<<<<Generics>>>>")
	lessons.Generics()
}
