package main

import (
	"fmt"
)

func main() {
	//Arrays

	//1. first way of declaring and assigning an array in go
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "apple"
	fruitArr[1] = "banana"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])

	//2. second way -shorthand- of declaring and assigning an array in go
	fruitSlice := [3]string{"apple", "orange", "pear"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[2])
	fmt.Println(fruitSlice[1:3])
}
