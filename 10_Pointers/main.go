package main

import (
	"fmt"
)

func main() {
	//Pointer to the memory address/location of a value that's in a variable
	a := 5
	b := &a
	fmt.Println(a, " - ", b)
	fmt.Printf("%T - %T\n", a, b)

	//use * to read value from address
	fmt.Println(a)
	fmt.Println(*&a)
	fmt.Println(*b)
	fmt.Println(b)
	fmt.Println(&b)
}
