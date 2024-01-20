package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(l1, r1 int) int {
	return l1 + r1
}

func greet() {
	fmt.Println("Hello from here")
}
func main() {
	greet()
	dod := double(4)
	fmt.Println("a double is ", dod)

	pls := add(dod, 3)
	fmt.Println("pls is ", pls)

	another := add(double(6), 2)
	fmt.Println("anoter is ", another)
}
