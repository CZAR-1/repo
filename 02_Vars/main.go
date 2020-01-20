//multiple ways to declare variables

package main

import "fmt"

//Global variables go here
//6. Shorthand assignments cannot be done outside of the function

func main() {
	//1. using var keyword, type explicitly indicated
	var name string = "Czar"
	fmt.Println(name)

	//2. using var keyword, type does not need to be explicitly indicated,
	//is inferred from the right hand side (int in this case)
	//to cast this var as a specific type, you need to explicitly define it.
	// (int32)
	var age int32 = 37

	//3. All variables declared must also be used.
	fmt.Println(name, " - ", age)

	//4. printf is an fmt function with verbs that assist in defining type
	// these verbs are defined under fmt package overview
	fmt.Printf("%T %T\n", name, age)

	//5. const works as it does in java, does not allow to change the
	// value of a var
	const isCool = true
	//isCool = false is not possible under const
	fmt.Println(isCool)

	//7. Shorthand
	lastName := "Marquez"
	//this will be automatically types as float64. To make it float32,
	//you will need to explicitly type it.
	size := 1.34

	fmt.Printf("%T, %T\n", lastName, size)

	//8. Shorthand 2: You can couple automatically typable vars in one line
	// it's much better practice to do this with related vars and like types,
	//however
	maiden, length, width, middle := "Linda", 8, 7.23, "Sr."

	fmt.Printf("%T, %T, %T, %T\n", maiden, length, width, middle)
}
