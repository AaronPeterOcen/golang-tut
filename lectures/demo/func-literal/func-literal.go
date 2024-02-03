package main

import "fmt"

// // func hellow() {
// // 	fmt.Printf("Hello, ")
// // 	world := func ()  {
// // 		fmt.Printf("World!\n")
// // 	}
// // 	world()
// // 	world()
// // }

// func customMsg(fn func(m string), msg string) {
// 	msg = strings.ToUpper(msg)
// 	fn(msg)
// }

// func surround() func(msg string) {
// 	return func(msg string) {
// 		fmt.Printf("%.*s\n", len(msg), "--------")
// 		fmt.Println(msg)
// 		fmt.Printf("%.*s\n", len(msg), "--------")
// 	}
// }

// func calcPrice(
// 	subtt float64,
// 	discfn func(subtt float64) float64) float64 {
// 		return subtt - (subtt * discountfn(subtt))
// }

func add(a, b int) int {
	return a + b
}

func compu(a, b int, ope func(a, b int) int) int {
	fmt.Printf("running ops with %v and %v\n", a, b)
	return ope(a, b)
}

func main() {
	// hellow()
	// customMsg(surround(), "hheeyyahh")

	// discount := 0.1
	// discountfn := func(subtt float64) float64 {
	// 	if subtt > 100 {
	// 		discount += 0.1
	// 	}
	// 	if discount > 0.3 {
	// 		discount = 0.3
	// 	}
	// 	return discount
	// }
	// totalP := calcPrice(20.2, discountfn)
	// fmt.Println(totalP)

	fmt.Println("2+2 =", compu(2, 2, add))
	fmt.Println("10 - 2 =", compu(10, 2, func(a, b int) int {
		return a - b
	}))
	fmt.Println("10 * 2 =", compu(10, 2, func(a, b int) int {
		return a * b
	}))
	fmt.Println("10 / 2 =", compu(10, 2, func(a, b int) int {
		return a / b
	}))
}
