package main

import "fmt"

func main() {
	var myName = "Ocen"
	fmt.Print("my name is", myName, "\n")

	// using type annotations
	var name string = "Aaron"
	fmt.Print("name = ", name, "\n")

	// using create and assign
	userName := "admin"

	fmt.Print(userName)

	var sum int
	fmt.Print("The sum is ", sum, "\n")

	one, two := 1, 2
	fmt.Print("one is ", one, " two is ", two, "\n")

	// ok error idiom
	oneone, two := 3, 0
	fmt.Print("one is ", oneone, " two is ", two, "\n")

	sum = one + oneone
	fmt.Print(sum)

	var (
		l1 = "vars"
		l2 = "lars"
	)
	fmt.Print("l1", l1, "\n")
	fmt.Print("l2", l2, "\n")

	// ignoring one of the variables using _
	w1, w2, _ := "hone", "wole", "lala"
	fmt.Print(w1, w2)

}
