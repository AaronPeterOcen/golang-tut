package main

import "fmt"

type Room struct {
	name string
	cleaned bool
}

func checkClean(room [4]Room) {
	for i := 0; i < len(room); i++ {
		rooms := room[i]
		if rooms.cleaned {
			fmt.Println(rooms.name,"is clean")
		} else {
			fmt.Println(rooms.name, "is not clean")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "WareHouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkClean(rooms)
	
	fmt.Println("Performing cleaning...")
	rooms[1].cleaned = true
	rooms[3].cleaned = true

	checkClean(rooms)

	// myArray := [...]int{7,4,3}
	// // var myArray [3]int

	// for i := 0; i < len(myArray); i++ {
	// 	item := myArray[i]
	// 	fmt.Println(item)
	// 	// fmt.Println(myArray[2])
	// }
}
