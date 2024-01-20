//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greets(name string) string {
	return name

}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func any() {
	fmt.Println("Lolz")
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func add(one, two, three int) int {
	return one + two + three
}

// * Write a function that returns any number
func anyOne(one int) int {
	return one
}

// * Write a function that returns any two numbers
func two() (int, int) {
	return 1, 2
}

//* Call every function at least once

func main() {

	names := greets("ocen")
	fmt.Println("How are you ", names)

	any()

	daa := add(1, 2, 4)
	fmt.Println(daa)

	aneone := anyOne(1)
	fmt.Println(aneone)

	a, b := two()
	fmt.Println(a, b)

	// * Add three numbers together using any combination of the existing functions.
	addThree := add(aneone, a, b)
	//   - Print the result
	fmt.Println(addThree)
}
