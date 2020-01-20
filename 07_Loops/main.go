package main

import (
	"fmt"
)

func main() {
	//long method
	i := 1
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	//short method
	for i := 10; i <= 20; i++ {
		fmt.Printf("Number %d\n", i)
	}

	//fizzbuzz
	//common interview question:
	//loop through 100, for every
	// Multiple of 3, print fizz
	// Multiple of 5, print buzz
	// Multiple of 3 && 5, print fizzbuzz
	// else, print the integer iteration fo the loop

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
