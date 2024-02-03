package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitized []rune

	capt := func(r rune) {
		capitized = append(capitized, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)
	}

	for i := 0; i < len(data); i++ {
		go capt(data[i])
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Captlized: %c\n", capitized)

}

// function that prints text with a goroutine
// func count(amount int) {
// 	for i := 1; i <= amount; i++ {
// 		time.Sleep(1000 * time.Millisecond)
// 		fmt.Println(i)
// 	}
// }

// text to be printed
// go count(10)
// fmt.Println("wait for go")
// time.Sleep(10000 * time.Millisecond)
// fmt.Println("end program")
