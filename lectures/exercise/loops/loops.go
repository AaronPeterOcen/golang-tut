//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	for x := 1; x <= 50; x++ {
		//* Print integers 1 to 50, except:
		if x%3 == 0 {
			//  - Print "Fizz" if the integer is divisible by 3
			fmt.Println("Fizz")
		}
		if x%5 == 0 {
			//  - Print "Buzz" if the integer is divisible by 5
			fmt.Println("Buzz")
		}
		if x%5 == 0 && x%3 == 0 {
			//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
			fmt.Println("Fizzbuzz", x)
		}
		if x%5 != 0 && x%3 != 0 {
			fmt.Println(x)
		}
	}
}
