package main

import (
	"fmt"
	"math"
)

// type Preper interface {
// 	PreperDish()
// }

// type Chicken string
// type Salad string

// func (c Chicken) PreperDish() {
// 	fmt.Println("Cook Chicken")
// }

// func (s Salad) PreperDish() {
// 	fmt.Println("Chop Salad")
// 	fmt.Println("add dressing")
// }

// func prepMutli(dishes []Preper) {
// 	fmt.Println("Prepare dishes:")
// 	for i := 0; i < len(dishes); i++ {
// 		dish := dishes[i]
// 		fmt.Println("--Dish: %v--\n", dish)
// 		dish.PreperDish()
// 	}
// 	fmt.Println()
// }

// func main() {
// 	dish := []Preper{Chicken("Wings"), Salad("leaves")}
// 	prepMutli(dish)
// }

type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 13.4},
		circle{radius: 3.4},
		circle{radius: 6.4},
		square{length: 8.4},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("___")
	}
}