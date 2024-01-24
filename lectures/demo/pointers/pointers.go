package main

import "fmt"

type Counter struct {
	hits int
}

func plus(count *Counter) {
	count.hits += 1
	fmt.Println("Count", count)
}

func replace(old *string, new string, count *Counter) {
	*old = new
	plus(count)
}

func main() {
	count := Counter{}

	hello := "Jello"
	world := "World"
	fmt.Println(hello, world)

	replace(&hello, "hi", &count) 
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go", &count)
	fmt.Println(phrase)


	// var pointer_name *Data_Type
	// var s *string
	
	// normal variable declaration
	var a = 45

	// Initialization of pointer s with 
	// memory address of variable a
	var s *int
	fmt.Println("s = ", s)
	s = &a

	fmt.Println("value a = ", a)
	fmt.Println("Address of a", &a)
	fmt.Println("Value of s ", s)
	fmt.Println("value stored in s ", *s)
	*s = 12
	fmt.Println("value stored in s is now ", *s)


}
