//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//  information again

package main

import (
	"fmt"
)
type ShoppingList struct {
	item string
	price int
}

func buy(shop [4]ShoppingList){
	var cost, totals int
	for i := 0; i < len(shop); i++ {
		shoppy := shop[i] 
		cost += shoppy.price
		
		//* Add a fourth product to the list and print out the
		if shoppy.item != "" {
			totals += 1
		}
	}
	//* Print to the terminal:
	
	//  - The last item on the list
	fmt.Println("Last item on list:", shop[totals-1])
	//  - The total number of items
	fmt.Println("total items", totals)
	//  - The total cost of the items
	fmt.Println("total cost", cost)
}


func main() {
	listItem := [4]ShoppingList {
		{"Mango", 3},
		{"Eggs", 12},
		{"veggies", 6},
	}

	buy(listItem)

	listItem[3] = ShoppingList{"Chicken", 13}
	buy(listItem)
}
