package main

import "fmt"

func priceSlice(title string, slice []string){
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}


func main() {
	route := []string{"Movies","Salon","Shop"}
	priceSlice("Route 1", route)

	route = append(route, "Bar")

	priceSlice("Route 2", route)

	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")

	route = route[2:]
	priceSlice("remaining locales", route)
}
