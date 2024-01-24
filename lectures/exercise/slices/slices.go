//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func contents(line []Part){
	for i := 0; i <len(line); i++ {
		parts := line[i]
		fmt.Println(parts)
	}
}

func main() {
	line := make([]Part, 3)

	line[0] = "one"
	line[1] = "two"
	line[2] = "three"

	fmt.Println("3")
	contents(line)

	line = append(line, "four", "five")
	fmt.Println("2")
	contents(line)

	line = line[3:]
	fmt.Println("4n5")
	contents(line)
}
