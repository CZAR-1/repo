//maps are key value pairs

package main

import (
	"fmt"
)

func main() {
	//Define map
	//map[*key value type*] *data value type*
	emails := make(map[string]string)
	//The make() function is a special built-in function
	// that is used to initialize slices, maps, and channels.

	//assign key values
	emails["Bob"] = "bobsaget@gmail.com"
	emails["Shanon"] = "sharvy1@gmail.com"
	emails["Joe"] = "joe1@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Joe"])
	delete(emails, "Bob")

	fmt.Println(emails)

	//initialize map with composite literal
	czarMap := map[string]bool{
		"Java":       true,
		"Go":         true,
		"C++":        true,
		"Javascript": true,
		"Rust":       false,
	}

	czarMap["Assembly"] = false

	fmt.Println(czarMap)
}
