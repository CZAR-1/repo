package main

import (
	"fmt"
)

func main() {
	x, y := 10, 5
	//IF/ELSEIF/ELSE ommon convention is to not use parenthesis
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
		fmt.Printf("%d is greater than %d\n", y, x)
	} else if x == y {
		fmt.Printf("%d is equivalent to %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//Switch
	color := "red"
	if color == "blue" {
		fmt.Println("color is blue")
	} else if color == "red" {
		fmt.Println("color is red")
	} else {
		fmt.Println("color is not red or blue")
	}

}
