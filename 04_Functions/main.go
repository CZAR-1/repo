package main

import (
	"fmt"
)

// func functionName(inputType(s)) outputType{}
func greeting(name string) string {
	return "Hello " + name + "!"
}

//(num1 int, num2 int) can be condensed to (num1, num2 int)
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Hello World")

	// greeting function does not output, it returns a value,
	// thus needs to be wrapped within the fmt.Println function
	fmt.Println(greeting("Czar"))
	fmt.Println(getSum(4, 3))
}
