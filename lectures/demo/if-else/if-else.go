package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	q1, q2, q3 := 6, 4, 7

	if q1 > q2 {
		fmt.Println("q1 is higher")
	} else if q1 < q2 {
		fmt.Println("q2 is higher")
	} else {
		fmt.Println("q1 and q2 are equal")
	}

	if average(q1, q2, q3) > 7 {
		fmt.Println("average grades")
	} else {
		fmt.Println("instructor need to up his game")
	}

}
