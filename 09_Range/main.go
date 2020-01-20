package main

import (
	"fmt"
)

func main() {
	ids := []int{234, 534, 256, 799, 345, 432, 875}

	//loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//index is not necessary, but cannot initialize without using
	// so _ is used in its place, example
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add all IDs together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("The sum of IDs is %d\n", sum)

	//Range with maps
	//initialize map with composite literal
	czarMap := map[string]bool{
		"Java":       true,
		"Go":         true,
		"C++":        true,
		"Javascript": true,
		"Rust":       false,
	}

	for k, v := range czarMap {
		fmt.Printf("%s: %t\n", k, v)
	}

	for k := range czarMap {
		fmt.Printf("Language Name: %s\n", k)
	}
}
