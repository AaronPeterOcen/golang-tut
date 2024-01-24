// // An Empty map
// map[Key_Type]Value_Type{}

// // Map with key-value pair
// map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

package main

import "fmt"

func main() {


	// Creating and initializing empty map
	// Using var keyword
	var map1 map[int] int 

	// map1 := {
	// 		i: 1,
	// 	}
	

	if map1 == nil {
		fmt.Println("true")
	} else {
		fmt.Println("False")
	}

	  // Creating and initializing a map
	// Using shorthand declaration and
	// using map literals
	map2 := map[int]string{
		91: "a",
		92: "b",
		93: "c",
		94: "d",
		95: "e",
	}
	fmt.Println("Map-2:", map2)


	// using make function
	// make(map[Key_Type]Value_Type, initial_Capacity)
	// make(map[Key_Type]Value_Type)

	// Go program to illustrate how to
	// create and initialize a map
	// Using make() function

	var Mymap = make(map[float64]string)
	fmt.Println(Mymap)


	// As we already know that make() function
	// always returns a map which is initialized
	// So, we can add values in it
	Mymap[1.3] = "ocen"
	Mymap[1.5] = "Summert"
	fmt.Println(Mymap)

	// Go program to illustrate how
	// to iterate the map using for
	// range loop


	// Creating and initializing a map
	m_a_p := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	// Iterating map using for range loop
	for id, pet := range m_a_p {

		fmt.Println(id, pet)
	}

	list := make(map[string]int)
	list["tea"] = 11
	list["milk"] = 10
	list["bread"]++
	list["tea"] += 11 
	fmt.Println(list)

	delete(list, "bread")

	fmt.Println(list)
	fmt.Println("want", list["milk"], "milk")

	cereal, found := list["cereal"]

	if !found {
		fmt.Println("none")
	} else {
		fmt.Println("yeah", cereal, "boxes")
	}

	totalcount := 0
	for _, amount := range list {
		totalcount += amount
	}
	fmt.Println("total count is", totalcount)
}
