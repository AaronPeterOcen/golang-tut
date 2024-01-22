package main

import (
	"fmt"
)

func main() {
	ia := 0
	fmt.Println("sum", ia)

	i := 0
	for i = ia; i <= 5; i++ {
		ia += i
		fmt.Println("sum is", ia)
	}

	for ia > 5 {
		ia -= 5
		fmt.Println("Dec. sum is", ia)
	}
}
